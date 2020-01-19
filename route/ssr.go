package route

import (
	"github.com/JabinGP/ssr-sub-hub/controller"
	"github.com/JabinGP/ssr-sub-hub/middleware"
	"github.com/kataras/iris/v12/core/router"
)

func routeSSR(party router.Party) {
	party.Get("/", middleware.Verify, controller.GetSSR)
	party.Get("/list", middleware.Verify, controller.GetSSRList)
}
