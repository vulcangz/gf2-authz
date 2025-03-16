package v1

import (
	"github.com/vulcangz/gf2-authz/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AuthReq struct {
	g.Meta   `path:"/auth" tags:"auth" method:"post" summary:"authentication"`
	Username string `p:"username" dc:"john.doe"`
	Password string `p:"password" dc:"mypassword"`
}

type AuthRes struct {
	AccessToken string       `json:"access_token"`
	TokenType   string       `json:"token_type"`
	ExpiresIn   int64        `json:"expires_in"`
	User        *entity.User `json:"user"`
}

type TokenReq struct {
	g.Meta       `path:"/token" tags:"auth" method:"post" summary:"authentication"`
	GrantType    string `p:"grant_type" dc:"client_credentials"`
	ClientID     string `p:"client_id" dc:"0be4e0e0-6788-4b99-8e00-e0af5b4945b1"`
	ClientSecret string `p:"client_secret" dc:"EXCAdNZjCz0qJ_8uYA2clkxVdp_f1tm7"`
	RefreshToken string `p:"refresh_token"`
}

// type TokenRes struct {
// 	g.Meta `mime:"application/json"`
// 	Data   *tokenRes `json:"data"`
// }

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
}
