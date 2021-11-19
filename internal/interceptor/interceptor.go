package interceptor

import (
	"context"

	"github.com/fatih/twirpdemo/internal/rpc"
	"github.com/twitchtv/twirp"
)

func NewInterceptor() twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if m, ok := twirp.MethodName(ctx); ok && m == "Add" {
				return next(ctx, &rpc.AddReq{A: 4, B: 3})
			}

			return next(ctx, req)
		}
	}
}
