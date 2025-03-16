package jwt

import (
	"time"
)

type Auth struct {
	AccessTokenDuration  time.Duration `config:"auth_access_token_duration"`
	RefreshTokenDuration time.Duration `config:"auth_refresh_token_duration"`
	Domain               string        `config:"auth_domain"`
	JWTSignString        []byte        `config:"auth_jwt_sign_string"`
}
