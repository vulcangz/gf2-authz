package jwt

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/model"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
)

var JWTManager Manager

const (
	TokenTypeBearer = "bearer"
)

type Claims struct {
	*model.ContextUser
	jwt.RegisteredClaims
}

type Token struct {
	Token     string
	TokenType string
	ExpiresIn int64
}

type Manager interface {
	Generate(identifier string) (*Token, error)
	Parse(accessToken string) (*Claims, error)
}

type manager struct {
	cfg   *entity.AuthConfig
	clock ctime.Clock
}

func NewManager(
	cfg *entity.AuthConfig,
	clock ctime.Clock,
) Manager {
	// Ensure JWT library is using our clock.
	jwt.TimeFunc = clock.Now

	return &manager{
		cfg:   cfg,
		clock: clock,
	}
}

func (m *manager) Generate(identifier string) (*Token, error) {
	now := m.clock.Now()
	ctx := gctx.New()
	expireAt := now.Add(m.cfg.AccessTokenDuration)
	appName := g.Cfg().MustGet(ctx, "app.name").String()
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    appName,
			Subject:   identifier,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.cfg.JWTSignString)
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:     accessToken,
		TokenType: TokenTypeBearer,
		ExpiresIn: int64(expireAt.Sub(m.clock.Now()).Seconds()),
	}, nil
}

func (m *manager) Parse(accessToken string) (*Claims, error) {
	claims := &Claims{}

	jwtToken, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return m.cfg.JWTSignString, nil
	})
	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}

func Init(ctx context.Context) {
	authCfg, _ := getAuthConfig(ctx)
	var clock ctime.Clock
	if gmode.IsTesting() {
		clock = ctime.NewStaticClock()
	} else {
		clock = ctime.NewClock()
	}
	JWTManager = NewManager(authCfg, clock)
}

// GetAuth get auth configuration options
func getAuthConfig(ctx context.Context) (conf *entity.AuthConfig, err error) {
	err = g.Cfg().MustGet(ctx, "auth").Scan(&conf)
	if conf == nil {
		model := &entity.AuthConfig{}
		conf = model.DefaultConfig()
	}
	return
}
