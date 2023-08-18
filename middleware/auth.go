package middleware

import (
	"context"
	"errors"
	"warehouse/helper"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				auth := tr.RequestHeader().Get("Authorization")

				if auth == "" {
					return nil, errors.New("Auth fail")
				}

				userClaims, err := helper.ParseToken(auth)

				if err != nil {
					return nil, err
				}

				if userClaims.Identity == "" {
					return nil, errors.New("Auth fail")
				}

				// 将认证中间件校验通过的解析信息写入到上下文meta元信息中，方便直接使用
				ctx = metadata.NewServerContext(ctx, metadata.New(map[string][]string{
					"username": {userClaims.Name},
					"identity": {userClaims.Identity},
				}))
			}
			return handler(ctx, req)
		}
	}
}
