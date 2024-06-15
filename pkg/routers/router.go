package routers

import (
	"github.com/labstack/echo/v4"
	"ussd-gateway/pkg/controllers"
	"ussd-gateway/pkg/di"
	"ussd-gateway/pkg/utils"
)

type RouterEngine interface {
	ApiRoutes(di.ServiceCollection)
}

type routerEngine struct {
	app *echo.Echo
}

func (engine routerEngine) ApiRoutes(services di.ServiceCollection) {
	gatewayRoute := engine.app.Group("/v1").Group("/ussd").Group("/gateways")
	gatewayController, _ := services.GetService(utils.GatewayController)
	gatewayHandler := gatewayController.(controllers.GatewayController)
	gatewayRoute.POST("/:gateway", gatewayHandler.Gateway)
}

func NewEngine(app *echo.Echo) RouterEngine {
	return &routerEngine{
		app: app,
	}
}
