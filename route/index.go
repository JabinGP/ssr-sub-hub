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
	v2ray := app.Party("/v2ray")
	{
		routeV2ray(v2ray)
	}
	user := app.Party("/user")
	{
		routeUser(user)
	}
	routeStatic(app)
}
