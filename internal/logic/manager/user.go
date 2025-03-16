package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/token"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type sUserManager struct {
	repository          orm.UserRepository
	principalRepository orm.PrincipalRepository
	transactionManager  orm.TransactionManager
	tokenGenerator      token.Generator
}

// NewUser initializes a new user manager.
func NewUser(
	repository orm.UserRepository,
	principalRepository orm.PrincipalRepository,
	transactionManager orm.TransactionManager,
	tokenGenerator token.Generator,
) *sUserManager {
	return &sUserManager{
		repository:          repository,
		principalRepository: principalRepository,
		transactionManager:  transactionManager,
		tokenGenerator:      tokenGenerator,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	userRepository := orm.New[entity.User](db)
	principalRepository := orm.New[entity.Principal](db)
	tm := orm.NewTransactionManager(db)
	tg := token.NewGenerator()

	service.RegisterUserManager(NewUser(userRepository, principalRepository, tm, tg))
}

func (m *sUserManager) GetRepository() orm.UserRepository {
	return m.repository
}

func (m *sUserManager) Create(ctx context.Context, username string, password string) (*entity.User, error) {
	exists, err := m.repository.GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: username},
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to check for existing user: %v", err)
	}

	if exists != nil {
		return nil, fmt.Errorf("a user already exists with username %q", username)
	}

	if password == "" {
		password, err = m.tokenGenerator.Generate(10)
		if err != nil {
			return nil, fmt.Errorf("unable to generate a random password: %v", err)
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	user := &entity.User{
		Username:     username,
		Password:     password,
		PasswordHash: string(hashedPassword),
	}

	if err := m.repository.WithTransaction(transaction).Create(user); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to create user: %v", err)
	}

	if err := m.principalRepository.WithTransaction(transaction).Create(&entity.Principal{
		ID: entity.UserPrincipal(user.Username),
	}); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to create user: %v", err)
	}

	return user, nil
}

func (m *sUserManager) Delete(ctx context.Context, username string) error {
	user, err := m.GetRepository().GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: username},
	})
	if err != nil {
		return fmt.Errorf("cannot retrieve user: %v", err)
	}

	// Retrieve principal
	principal, err := m.principalRepository.Get(
		entity.UserPrincipal(username),
	)
	if err != nil {
		return fmt.Errorf("cannot retrieve user principal: %v", err)
	}

	// Delete both user and principal
	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	if err := m.principalRepository.WithTransaction(transaction).Delete(principal); err != nil {
		_ = transaction.Rollback()
		return fmt.Errorf("cannot delete principal: %v", err)
	}

	if err := m.GetRepository().WithTransaction(transaction).Delete(user); err != nil {
		_ = transaction.Rollback()
		return fmt.Errorf("cannot delete user: %v", err)
	}

	return nil
}

func (m *sUserManager) UpdatePassword(username string, password string) error {
	user, err := m.repository.GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: username},
	})
	if err != nil {
		return fmt.Errorf("unable to retrieve user: %v", err)
	}

	if password == "" {
		password, err = m.tokenGenerator.Generate(10)
		if err != nil {
			return fmt.Errorf("unable to generate a random password: %v", err)
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = password
	user.PasswordHash = string(hashedPassword)

	if err := m.repository.Update(user); err != nil {
		return fmt.Errorf("unable to update user password: %v", err)
	}

	return nil
}
