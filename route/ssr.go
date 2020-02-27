package route

import (
	"github.com/JabinGP/ssr-sub-hub/controller"
	"github.com/kataras/iris/v12/core/router"
)

func routeSSR(party router.Party) {
	party.Get("/{username:string}/{password:string}", controller.GetSSR)
	party.Get("/list/{username:string}/{password:string}", controller.GetSSRList)
}
