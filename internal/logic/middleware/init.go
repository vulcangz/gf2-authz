package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/vulcangz/gf2-authz/internal/service"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
