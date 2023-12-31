package server

import (
	"warehouse/api/git"
	v1 "warehouse/api/helloworld/v1"
	"warehouse/internal/conf"
	"warehouse/internal/service"
	"warehouse/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(middleware.Auth()).Path("/api.git.User/Login").Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)

	// 注册用户相关service
	git.RegisterUserHTTPServer(srv, service.NewUserService())
	// 注册仓库模块服务
	git.RegisterRepoHTTPServer(srv, service.NewRepoService())
	return srv
}
