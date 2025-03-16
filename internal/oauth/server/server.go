package server

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/gogf/gf/v2/frame/g"
)

func NewServer(
	manager oauth2.Manager,
) *server.Server {
	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		g.Log().Printf(context.Background(), "Response Error: %#+v\n", re)
	})

	return srv
}
