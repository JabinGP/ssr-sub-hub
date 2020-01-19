package middleware

import (
	"github.com/kataras/iris/v12/context"
)

// Verify Verify Handler
var Verify context.Handler

func initVerify() {
	Verify = func(ctx context.Context) {
		ctx.Application().Logger().Println("using verify")
		ctx.Next()
	}
}
