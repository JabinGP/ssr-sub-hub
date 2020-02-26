package route

import (
	"github.com/JabinGP/ssr-sub-hub/controller"
	"github.com/kataras/iris/v12/core/router"
)

func routeV2ray(party router.Party) {
	party.Get("/", controller.GetV2ray)
	party.Get("/list", controller.GetV2rayList)
}
