package route

import (
	"github.com/JabinGP/ssr-sub-hub/controller"
	"github.com/kataras/iris/v12/core/router"
)

func routeUser(party router.Party) {
	party.Get("/config/{username:string}/{password:string}", controller.GetUserConfig)
	party.Post("/config", controller.PostUserConfig)
}
