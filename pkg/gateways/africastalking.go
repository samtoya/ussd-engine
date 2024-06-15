package gateways

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"ussd-gateway/pkg/contracts"
	"ussd-gateway/pkg/data/dtos"
	"ussd-gateway/pkg/di"
	"ussd-gateway/pkg/utils"
)

type africasTalkingGateway struct {
	di di.ServiceCollection
}

func (gtw africasTalkingGateway) HandleRequest(req interface{}) string {
	log.Println("----- Using Africastalking gateway engine -----")
	ctx := req.(echo.Context)
	dto := new(dtos.AfricasTalkingDto)
	if err := ctx.Bind(dto); err != nil {
		log.Println("Binding error:", err)
		return "failed to bind request"
	}

	v, _ := gtw.di.GetService(utils.Validator)
	if err := v.(*validator.Validate).Struct(dto); err != nil {
		log.Println("Binding error:", err)
		return "failed to validate request"
	}

	var response string
	if dto.Text == "" {
		// Lookup the initial screen from the yaml screens and route to that screen
		// Set the cache to the screenName of the initial screen
	}

	log.Println("AT Request Params:", dto)

	return response
}

func NewAfricasTalkingGateway(deps di.ServiceCollection) contracts.GatewayEngine {
	return &africasTalkingGateway{di: deps}
}
