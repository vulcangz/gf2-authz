package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/vulcangz/gf2-authz/internal/consts"
	"github.com/vulcangz/gf2-authz/internal/controller/api/action"
	"github.com/vulcangz/gf2-authz/internal/controller/api/audit"
	"github.com/vulcangz/gf2-authz/internal/controller/api/check"
	"github.com/vulcangz/gf2-authz/internal/controller/api/client"
	"github.com/vulcangz/gf2-authz/internal/controller/api/common"
	"github.com/vulcangz/gf2-authz/internal/controller/api/compiled"
	"github.com/vulcangz/gf2-authz/internal/controller/api/policy"
	"github.com/vulcangz/gf2-authz/internal/controller/api/principal"
	"github.com/vulcangz/gf2-authz/internal/controller/api/resource"
	"github.com/vulcangz/gf2-authz/internal/controller/api/role"
	"github.com/vulcangz/gf2-authz/internal/controller/api/stats"
	"github.com/vulcangz/gf2-authz/internal/controller/api/user"
	"github.com/vulcangz/gf2-authz/internal/lib/database"
	"github.com/vulcangz/gf2-authz/internal/lib/jwt"
	"github.com/vulcangz/gf2-authz/internal/service"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Usage: "http",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// s.SetServerRoot("resource/public")
			sCfg, _ := service.SysConfig().GetHTTPServer(ctx)
			if sCfg != nil {
				s.SetConfigWithMap(g.Map{
					"Address":    sCfg.Address,
					"ServerRoot": sCfg.ServerRoot,
				})
			}

			s.Use(
				service.Middleware().MiddlewareCORS,
				ghttp.MiddlewareHandlerResponse)

			// Remove all builtin metrics that are produced by prometheus client.
			prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
			prometheus.Unregister(collectors.NewGoCollector())
			handler := promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})
			s.BindHandler("/v1/metrics", ghttp.WrapH(handler))

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.POST("/auth", common.Auth.Authenticate)
					group.GET("/oauth", common.OAuthAuthenticate)
					group.GET("/oauth/callback", common.OAuthCallback)
					group.POST("/token", common.Auth.TokenNew)

					// Authz resources
					group.Middleware(service.Middleware().Authentication)

					group.POST("/check", check.Check)

					group.Middleware(service.Middleware().Authorization)

					group.Group("/actions", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /":            action.Action.List,
							"GET: /:identifier": action.Action.Get,
						})
					})

					group.Group("/audits", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /": audit.Audit.AuditGet,
						})
					})

					group.Group("/clients", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /":               client.Client.List,
							"POST: /":              client.Client.Create,
							"GET: /:identifier":    client.Client.Get,
							"Delete: /:identifier": client.Client.Delete,
						})
					})

					group.Group("/compiled", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /": compiled.Compiled.List,
						})
					})

					group.Group("/policies", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /":               policy.Policy.List,
							"POST: /":              policy.Policy.Create,
							"GET: /:identifier":    policy.Policy.Get,
							"Delete: /:identifier": policy.Policy.Delete,
							"Put: /:identifier":    policy.Policy.Update,
						})
					})

					group.Group("/principals", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /":               principal.Principal.List,
							"POST: /":              principal.Principal.Create,
							"GET: /:identifier":    principal.Principal.Get,
							"Delete: /:identifier": principal.Principal.Delete,
							"Put: /:identifier":    principal.Principal.Update,
						})
					})

					group.Group("/resources", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /":               resource.Resource.List,
							"POST: /":              resource.Resource.Create,
							"GET: /:identifier":    resource.Resource.Get,
							"Delete: /:identifier": resource.Resource.Delete,
							"Put: /:identifier":    resource.Resource.Update,
						})
					})

					group.Group("/roles", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /":               role.Role.List,
							"POST: /":              role.Role.Create,
							"GET: /:identifier":    role.Role.Get,
							"Delete: /:identifier": role.Role.Delete,
							"Put: /:identifier":    role.Role.Update,
						})
					})

					group.Group("/stats", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /": stats.Stats.Get,
						})
					})

					group.Group("/users", func(group *ghttp.RouterGroup) {
						group.Hook("/*any", ghttp.HookBeforeServe, hookForAuthorization)
						group.Map(g.Map{
							"GET: /":               user.User.List,
							"POST: /":              user.User.Create,
							"GET: /:identifier":    user.User.Get,
							"Delete: /:identifier": user.User.Delete,
						})
					})

				})
			})

			initComponents(ctx)
			s.Run()
			return nil
		},
	}
)

func initComponents(ctx context.Context) {
	// database connection initialize
	database.GetDatabase(ctx)

	jwt.Init(ctx)

	common.Initializer(ctx)
}

func hookForAuthorization(r *ghttp.Request) {
	action := gstr.ToLower(r.Method)

	if action != "options" {
		seg := ""
		act := ""
		namedPath := false

		path := r.URL.Path
		if parr := gstr.Split(path, "/"); len(parr) > 2 {
			seg = parr[2]

			if len(parr) > 3 {
				namedPath = true
			}
		}

		if namedPath {
			if action == "put" {
				action = "update"
			}
			act = action
		} else {
			switch action {
			case "get":
				act = "list"
				// expect audits & stats
				if seg == "audits" || seg == "stats" {
					act = "get"
				}
			case "post":
				act = "create"
			default:
				act = action
			}
		}
		k := gstr.ToLower(seg + "-" + act)
		r.SetCtxVar(consts.ResourceKindKey, ResourcesAndActionsByMethod[k][0])
		r.SetCtxVar(consts.ResourceValueKey, r.GetParam("identifier", "*").String())
		r.SetCtxVar(consts.ActionKey, ResourcesAndActionsByMethod[k][1])
	}
}
