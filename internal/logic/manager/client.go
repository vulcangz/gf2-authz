package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	"github.com/vulcangz/gf2-authz/internal/lib/token"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

type sClientManager struct {
	repository          orm.ClientRepository
	principalRepository orm.PrincipalRepository
	transactionManager  orm.TransactionManager
	tokenGenerator      token.Generator
}

// NewClient initializes a new client manager.
func NewClient(
	repository orm.ClientRepository,
	principalRepository orm.PrincipalRepository,
	transactionManager orm.TransactionManager,
	tokenGenerator token.Generator,
) *sClientManager {
	return &sClientManager{
		repository:          repository,
		principalRepository: principalRepository,
		transactionManager:  transactionManager,
		tokenGenerator:      tokenGenerator,
	}
}

func init() {
	db := database.GetDatabase(context.Background())
	clientRepository := orm.New[entity.Client](db)
	principalRepository := orm.New[entity.Principal](db)
	tm := orm.NewTransactionManager(db)
	tg := token.NewGenerator()
	service.RegisterClientManager(NewClient(clientRepository, principalRepository, tm, tg))
}

func (m *sClientManager) GetRepository() orm.ClientRepository {
	return m.repository
}

func (m *sClientManager) Create(ctx context.Context, name string, domain string) (*entity.Client, error) {
	exists, err := m.repository.GetByFields(map[string]orm.FieldValue{
		"name": {Operator: "=", Value: name},
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("unable to check for existing client: %v", err)
	}

	if exists != nil {
		return nil, fmt.Errorf("a client already exists with name %q", name)
	}

	clientID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("unable to generate client identifier: %v", err)
	}

	secret, err := m.tokenGenerator.Generate(48)
	if err != nil {
		return nil, fmt.Errorf("unable to generate client secret: %v", err)
	}

	client := &entity.Client{
		ID:     clientID.String(),
		Secret: secret,
		Domain: domain,
		Name:   name,
	}

	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	if err := m.repository.WithTransaction(transaction).Create(client); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to create client: %v", err)
	}

	if err := m.principalRepository.WithTransaction(transaction).Create(&entity.Principal{
		ID:       entity.ClientPrincipal(client.Name),
		IsLocked: true,
	}); err != nil {
		_ = transaction.Rollback()
		return nil, fmt.Errorf("unable to create principal: %v", err)
	}

	return client, nil
}

func (m *sClientManager) Delete(ctx context.Context, identifier string) error {
	client, err := m.GetRepository().Get(identifier)
	if err != nil {
		return fmt.Errorf("cannot retrieve client: %v", err)
	}

	// Retrieve principal
	principal, err := m.principalRepository.Get(
		entity.ClientPrincipal(client.Name),
	)
	if err != nil {
		return fmt.Errorf("cannot retrieve client principal: %v", err)
	}

	// Delete both client and principal
	transaction := m.transactionManager.New()
	defer func() { _ = transaction.Commit() }()

	if err := m.principalRepository.WithTransaction(transaction).Delete(principal); err != nil {
		_ = transaction.Rollback()
		return fmt.Errorf("cannot delete principal: %v", err)
	}

	if err := m.GetRepository().WithTransaction(transaction).Delete(client); err != nil {
		_ = transaction.Rollback()
		return fmt.Errorf("cannot delete client: %v", err)
	}

	return nil
}
