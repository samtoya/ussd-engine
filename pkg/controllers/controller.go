package controllers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"ussd-gateway/pkg/contracts"
	"ussd-gateway/pkg/di"
	"ussd-gateway/pkg/utils"
)

type GatewayController interface {
	Gateway(echo.Context) error
}

type gatewayController struct {
	di di.ServiceCollection
}

func (controller *gatewayController) Gateway(ctx echo.Context) error {
	gateway := ctx.Param("gateway")
	log.Printf("----- Service Entrypoint: %s ----- \n", gateway)
	service, _ := controller.di.GetService(utils.GatewayFactory)
	gatewayFactory := service.(contracts.GatewayFactory)
	gatewayEngine := gatewayFactory.Detect(gateway)
	response := gatewayEngine.HandleRequest(ctx)
	return ctx.String(http.StatusOK, response)
}

func NewGatewayController(d di.ServiceCollection) GatewayController {
	return &gatewayController{
		di: d,
	}
}
