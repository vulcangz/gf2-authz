package jwt

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

func TestNewManager(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	cfg := &entity.AuthConfig{}
	clock := ctime.NewMockClock(ctrl)

	// When
	managerInstance := NewManager(cfg, clock)

	// Then
	assert.IsType(t, new(manager), managerInstance)
}

func TestManager_Generate(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	cfg := &entity.AuthConfig{
		AccessTokenDuration: 1 * time.Hour,
	}

	date := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	clock := ctime.NewMockClock(ctrl)
	clock.EXPECT().Now().Return(date).Times(2)

	manager := NewManager(cfg, clock)

	// When
	token, err := manager.Generate("user-123")

	// Then
	assert.Nil(t, err)
	assert.Equal(t, &Token{
		Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoeiIsInN1YiI6InVzZXItMTIzIiwiZXhwIjoxNjcyNTM0ODAwLCJpYXQiOjE2NzI1MzEyMDB9.v4KOq5Fw-2PoVPNNMDgunVO4R_RO0NwdLUlyupNXvSk",
		TokenType: "bearer",
		ExpiresIn: 3600,
	}, token)
}

func TestManager_Parse_WhenInvalidToken(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	cfg := &entity.AuthConfig{
		AccessTokenDuration: 1 * time.Hour,
	}

	date := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	clock := ctime.NewMockClock(ctrl)
	clock.EXPECT().Now().Return(date).Times(3)

	manager := NewManager(cfg, clock)

	// When
	token, err := manager.Generate("user-123")
	assert.Nil(t, err)

	claims, err := manager.Parse(token.Token)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "user-123", claims.Subject)
}

func TestManager_Parse_WhenSuccess(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)

	cfg := &entity.AuthConfig{
		AccessTokenDuration: 1 * time.Hour,
	}

	clock := ctime.NewMockClock(ctrl)

	manager := NewManager(cfg, clock)

	// When
	claims, err := manager.Parse("this-is-an-invalid-token")

	// Then
	assert.Equal(t, "token contains an invalid number of segments", err.Error())
	assert.Nil(t, claims)
}
