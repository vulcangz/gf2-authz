package handler

import (
	"context"
	"errors"

	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var (
	// ClientNotFoundErr is returned when client_id does not exists.
	ClientNotFoundErr = status.Error(codes.NotFound, "client not found")

	// InvalidCredentialsErr is returned when client_id or client_secret is invalid.
	InvalidCredentialsErr = status.Error(codes.InvalidArgument, "invalid credentials")
)

type Auth interface {
	Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error)
}

type auth struct {
	tokenManager jwt.Manager
}

func NewAuth(
	tokenManager jwt.Manager,
) Auth {
	return &auth{
		tokenManager: tokenManager,
	}
}

func (h *auth) Authenticate(ctx context.Context, req *v1.AuthenticateRequest) (*v1.AuthenticateResponse, error) {
	client, err := service.ClientManager().GetRepository().Get(req.GetClientId())
	if err != nil || client.Secret != req.GetClientSecret() {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ClientNotFoundErr
		}

		return nil, InvalidCredentialsErr
	}

	token, err := h.tokenManager.Generate(entity.ClientPrincipal(client.Name))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.AuthenticateResponse{
		Token:     token.Token,
		Type:      token.TokenType,
		ExpiresIn: token.ExpiresIn,
	}, nil
}
