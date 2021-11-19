package hook

import (
	"context"
	"log"

	"github.com/twitchtv/twirp"
)

func NewErrorHook() *twirp.ServerHooks {
	hooks := &twirp.ServerHooks{
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			log.Println("received an error", twerr)
			return ctx
		},
	}

	return hooks

}
