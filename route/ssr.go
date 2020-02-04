package route

import (
	"github.com/JabinGP/ssr-sub-hub/controller"
	"github.com/kataras/iris/v12/core/router"
)

func routeSSR(party router.Party) {
	party.Get("/", controller.GetSSR)
	party.Get("/list", controller.GetSSRList)
}
