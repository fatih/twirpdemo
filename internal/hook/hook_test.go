package hook

import (
	"context"
	"errors"
	"testing"

	"github.com/twirp-ecosystem/twirptest"
	"github.com/twitchtv/twirp"
)

func TestNewErrorHook(t *testing.T) {
	errorHook := NewErrorHook()

	isInvoked := false

	testHook := &twirp.ServerHooks{
		Error: func(ctx context.Context, twerr twirp.Error) context.Context {
			isInvoked = true

			ctx = errorHook.Error(context.WithValue(ctx, "foo", "bar"), twerr)
			val := ctx.Value("foo")
			if val != "bar" {
				t.Errorf("want: %s, got: %s", "bar", val)
			}

			return ctx
		},
	}

	server, client := twirptest.ServerAndClient(
		twirptest.ErroringHatmaker(errors.New("sorry, I'm closed")),
		testHook,
	)
	t.Cleanup(func() { server.Close() })

	_, err := client.MakeHat(context.Background(), &twirptest.Size{Inches: 1})
	if err == nil {
		t.Error("expecting an error, got nil")
	}

	if !isInvoked {
		t.Error("error hook should be invoked")
	}

}
