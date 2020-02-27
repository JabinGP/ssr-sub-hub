package route

import (
	"github.com/JabinGP/ssr-sub-hub/controller"
	"github.com/kataras/iris/v12/core/router"
)

func routeV2ray(party router.Party) {
	party.Get("/{username:string}/{password:string}", controller.GetV2ray)
	party.Get("/list/{username:string}/{password:string}", controller.GetV2rayList)
}
