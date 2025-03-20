package common

import (
	"context"
	"errors"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/vulcangz/gf2-authz/api/api/common/v1"
	"github.com/vulcangz/gf2-authz/internal/lib/orm"
	response "github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/service"
	"golang.org/x/crypto/bcrypt"
)

var (
	Auth = cAuth{}

	clientErr       = gerror.NewCode(gcode.New(60001, "客户端错误", nil))
	accessTokenErr  = gerror.NewCode(gcode.New(60002, "access_token错误", nil))
	refreshTokenErr = gerror.NewCode(gcode.New(60003, "refresh_token错误", nil))
)

type cAuth struct {
}

// Authenticates a user
//
//	@security	Authentication
//	@Summary	Authenticates a user
//	@Tags		Auth
//	@Produce	json
//	@Param		default	body		AuthRequest	true	"Authentication request"
//	@Success	200		{object}	AuthResponse
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/auth [Post]
func (c *cAuth) Authenticate(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	user, err := service.UserManager().GetRepository().GetByFields(map[string]orm.FieldValue{
		"username": {Operator: "=", Value: req.Username},
	})
	if err != nil {
		response.ReturnError(r, http.StatusBadRequest, err, "")
		return
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	); err != nil {
		response.ReturnError(r, http.StatusBadRequest, err, "invalid credentials")
		return
	}

	token, err := jwtManager.Generate(user.Username)
	if err != nil {
		response.ReturnError(r, http.StatusBadRequest, err, "unable to generate token")
		return
	}

	res = &v1.AuthRes{
		AccessToken: token.Token,
		TokenType:   token.TokenType,
		ExpiresIn:   token.ExpiresIn,
		User:        user,
	}

	return
}

// Retrieve a client token
//
//	@security	Authentication
//	@Summary	Retrieve a client token
//	@Tags		Auth
//	@Produce	json
//	@Param		default	body		TokenRequest	true	"Token request"
//	@Success	200		{object}	TokenResponse
//	@Failure	400		{object}	model.ErrorResponse
//	@Failure	500		{object}	model.ErrorResponse
//	@Router		/v1/token [Post]
func (c *cAuth) TokenNew(ctx context.Context, req *v1.TokenReq) (res *v1.TokenRes, err error) {
	r := g.RequestFromCtx(ctx)
	w := r.Response.RawWriter()
	rr := r.Request
	res = new(v1.TokenRes)

	srv.SetClientInfoHandler(clientInfoHandler)
	srv.SetResponseTokenHandler(func(w http.ResponseWriter, data map[string]any, header http.Header, statusCode ...int) error {
		status := http.StatusOK
		if len(statusCode) > 0 && statusCode[0] > 0 {
			status = statusCode[0]
		}

		if status >= 200 && status < 300 {
			if v, ok := data["access_token"]; ok {
				res.AccessToken = v.(string)

			}
			if v, ok := data["expires_in"]; ok {
				res.ExpiresIn = gconv.Int(v)
			}
			if v, ok := data["token_type"]; ok {
				res.TokenType = v.(string)
			}
			if v, ok := data["refresh_token"]; ok {
				res.RefreshToken = v.(string)
			}

			return nil
		}

		return errors.New("unable to handle oauth request")
	})

	// 允许密码模式、刷新Token
	// srv.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.Refreshing)
	err = srv.HandleTokenRequest(w, rr)
	if err != nil {
		response.ReturnHTTPError(w, http.StatusInternalServerError, err)
	}

	r.Middleware.Next()
	return
}

func clientInfoHandler(r *http.Request) (clientID, clientSecret string, err error) {
	_ = r.ParseMultipartForm(0)
	_ = r.ParseForm()
	if r.Form.Get("grant_type") == "refresh_token" {
		ti, err := srv.Manager.LoadRefreshToken(r.Context(), r.Form.Get("refresh_token"))
		if err != nil {
			return "", "", refreshTokenErr
		}
		clientID = ti.GetClientID()
		if clientID == "" {
			return "", "", clientErr
		}
		cli, err := srv.Manager.GetClient(r.Context(), clientID)
		if err != nil {
			return "", "", clientErr
		}
		clientSecret = cli.GetSecret()
		if clientSecret == "" {
			return "", "", clientErr
		}
		return clientID, clientSecret, nil
	}
	clientID = r.Form.Get("client_id")
	if clientID == "" {
		return "", "", clientErr
	}
	clientSecret = r.Form.Get("client_secret")
	if clientSecret == "" {
		return "", "", clientErr
	}
	return clientID, clientSecret, nil

}
