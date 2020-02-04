package route

import (
	"github.com/kataras/iris/v12"
)

// Route ...
func Route(app *iris.Application) {
	ssr := app.Party("/ssr")
	{
		routeSSR(ssr)
	}
	user := app.Party("/user")
	{
		routeUser(user)
	}
	routeStatic(app)
}
