package common

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gogf/gf/v2/net/ghttp"
	response "github.com/vulcangz/gf2-authz/internal/lib/response"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"github.com/vulcangz/gf2-authz/internal/service"
	"gorm.io/gorm"
)

const (
	OAuthClaimEmailKey = "email"
	OAuthClaimNameKey  = "name"

	OAuthStateCookieName     = "authz_state"
	OAuthExpiresInCookieName = "authz_expires_in"
	OAuthTokenCookieName     = "authz_access_token"
	OAuthNonceCookieName     = "authz_nonce"
)

// Authenticates a user using an OAuth OpenID Connect provider
//
//	@security	Authentication
//	@Summary	Authenticates a user using an OAuth OpenID Connect provider
//	@Tags		Auth
//	@Success	302
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/oauth [Get]
func OAuthAuthenticate(r *ghttp.Request) {
	state, err := tokenGenerator.Generate(16)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, fmt.Errorf("unable to generate state: %v", err))
	}

	nonce, err := tokenGenerator.Generate(16)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, fmt.Errorf("unable to generate none: %v", err))
	}

	setCallbackCookie(r, OAuthStateCookieName, state, oauthClientManager.GetCookiesDomainName())
	setCallbackCookie(r, OAuthNonceCookieName, nonce, oauthClientManager.GetCookiesDomainName())

	r.Response.RedirectTo(
		oauthClientManager.GetConfig().AuthCodeURL(state, oidc.Nonce(nonce)),
		http.StatusFound,
	)
}

// Callback of the OAuth OpenID Connect provider authentication
//
//	@security	Authentication
//	@Summary	Callback of the OAuth OpenID Connect provider authentication
//	@Tags		Auth
//	@Success	200	{object}	AuthResponse
//	@Failure	400	{object}	model.ErrorResponse
//	@Failure	500	{object}	model.ErrorResponse
//	@Router		/v1/oauth/callback [Get]
func OAuthCallback(r *ghttp.Request) {
	ctx := r.Context()

	state := r.Cookie.Get(OAuthStateCookieName).String()

	if r.Get("state").String() != state {
		response.ReturnError(r, http.StatusBadRequest, errors.New("state did not match"))
	}

	oauth2Token, err := oauthClientManager.GetConfig().Exchange(ctx, r.Get("code").String())
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, fmt.Errorf("failed to exchange token: %v", err))
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		response.ReturnError(r, http.StatusInternalServerError, errors.New("no id_token field in oauth2 token"))
	}

	idToken, err := oauthClientManager.GetVerifier().Verify(ctx, rawIDToken)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, fmt.Errorf("failed to verify id token: %v", err))
	}

	nonce := r.Cookie.Get(OAuthNonceCookieName).String()

	if idToken.Nonce != nonce {
		response.ReturnError(r, http.StatusBadRequest, fmt.Errorf("nonce did not match: %v", err))
	}

	// Obtain user claims from OpenID Connect ID token.
	idTokenClaims := map[string]any{}

	if err := idToken.Claims(&idTokenClaims); err != nil {
		response.ReturnError(r, http.StatusInternalServerError, err)
	}

	emailValue, err := retrieveClaim(idTokenClaims, OAuthClaimEmailKey)
	if err != nil {
		response.ReturnError(r, http.StatusBadRequest, err)
	}

	nameValue, err := retrieveClaim(idTokenClaims, OAuthClaimNameKey)
	if err != nil {
		response.ReturnError(r, http.StatusBadRequest, err)
	}

	// Retrieve or create principal from user email.
	_, err = service.PrincipalManager().GetRepository().Get(
		entity.UserPrincipal(emailValue),
	)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if _, err = service.PrincipalManager().Create(
			ctx,
			entity.UserPrincipal(emailValue),
			[]string{},
			map[string]any{
				"name": nameValue,
			},
		); err != nil {
			response.ReturnError(r, http.StatusInternalServerError, fmt.Errorf("unable to create principal: %v", err))
		}
	} else if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, fmt.Errorf("unable to retrieve principal: %v", err))
	}

	// Generate access token.
	jwtToken, err := jwtManager.Generate(emailValue)
	if err != nil {
		response.ReturnError(r, http.StatusInternalServerError, fmt.Errorf("unable to generate jwt token: %v", err))
	}

	setCallbackCookie(r, OAuthExpiresInCookieName, strconv.FormatInt(jwtToken.ExpiresIn, 10), oauthClientManager.GetCookiesDomainName())
	setCallbackCookie(r, OAuthTokenCookieName, jwtToken.Token, oauthClientManager.GetCookiesDomainName())

	r.Response.RedirectTo(oauthClientManager.GetFrontendRedirectURL(), http.StatusFound)
}

func retrieveClaim(claims map[string]any, key string) (string, error) {
	value, ok := claims[key]
	if !ok {
		return "", fmt.Errorf("unable to retrieve claim from issuer: %s", key)
	}

	return value.(string), nil
}

func setCallbackCookie(c *ghttp.Request, name, value, domain string) {
	c.Cookie.SetHttpCookie(&http.Cookie{
		Domain:   domain,
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		HttpOnly: false,
	})
}
