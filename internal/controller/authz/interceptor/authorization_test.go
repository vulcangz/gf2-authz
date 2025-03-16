package interceptor

/*
import (
	"context"
	"errors"
	"testing"

	lib_jwt "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/service"
)

func TestAuthorizationFunc_WhenIsAllowed(t *testing.T) {
	// Given
	// ctrl := gomock.NewController(t)
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.ClaimsKey, &jwt.Claims{
		RegisteredClaims: lib_jwt.RegisteredClaims{
			Subject: "my-subject",
		},
	})

	resourceKind := "my-resource-kind"
	resourceValue := "my-resource-value"
	action := "my-action"

	compiledManager := service.MockCompiledPolicy()
	compiledManager.EXPECT().IsAllowed("my-subject", resourceKind, resourceValue, action).
		Return(true, nil)

	// When
	result := AuthorizationFunc()(ctx, resourceKind, resourceValue, action)

	// Then
	assert.Equal(t, true, result)
}

func TestAuthorizationFunc_WhenIsNotAllowed(t *testing.T) {
	// Given
	// ctrl := gomock.NewController(t)
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.ClaimsKey, &jwt.Claims{
		RegisteredClaims: lib_jwt.RegisteredClaims{
			Subject: "my-subject",
		},
	})

	resourceKind := "my-resource-kind"
	resourceValue := "my-resource-value"
	action := "my-action"

	compiledManager := service.MockCompiledPolicy()
	compiledManager.EXPECT().IsAllowed("my-subject", resourceKind, resourceValue, action).
		Return(false, nil)

	// When
	result := AuthorizationFunc()(ctx, resourceKind, resourceValue, action)

	// Then
	assert.Equal(t, false, result)
}

func TestAuthorizationFunc_WhenError(t *testing.T) {
	// Given
	// ctrl := gomock.NewController(t)
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.ClaimsKey, &jwt.Claims{
		RegisteredClaims: lib_jwt.RegisteredClaims{
			Subject: "my-subject",
		},
	})

	resourceKind := "my-resource-kind"
	resourceValue := "my-resource-value"
	action := "my-action"

	expectedErr := errors.New("this is an error returned by compiledManager.IsAllowed()")

	compiledManager := service.MockCompiledPolicy()
	compiledManager.EXPECT().IsAllowed("my-subject", resourceKind, resourceValue, action).
		Return(true, expectedErr)

	// When
	result := AuthorizationFunc()(ctx, resourceKind, resourceValue, action)

	// Then
	assert.Equal(t, false, result)
}

*/
