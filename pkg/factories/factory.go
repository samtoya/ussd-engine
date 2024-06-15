package factories

import (
	"log"
	"os"
	"strings"
	"ussd-gateway/pkg/contracts"
	"ussd-gateway/pkg/di"
	"ussd-gateway/pkg/utils"
)

type gatewayFactory struct {
	di di.ServiceCollection
}

func (factory *gatewayFactory) Detect(s string) contracts.GatewayEngine {
	log.Println("Detecting engine's existence = ", s)
	var service contracts.GatewayEngine
	switch strings.ToLower(s) {
	case utils.ArkesselGateway:
		engine, ok := factory.di.GetService(utils.ArkesselGateway)
		if !ok {
			log.Fatalf("%s is not a gateway engine", s)
		}
		service = engine.(contracts.GatewayEngine)
	case utils.AfricasTalkingGateway:
		engine, ok := factory.di.GetService(utils.AfricasTalkingGateway)
		if !ok {
			log.Fatalf("%s is not a gateway engine", s)
		}
		service = engine.(contracts.GatewayEngine)
	default:
		log.Printf("%s gateway engine has not yet been configured\n", s)
		os.Exit(2)
	}

	log.Println("Found gateway engine: ", s)
	return service
}

func New(collection di.ServiceCollection) contracts.GatewayFactory {
	return &gatewayFactory{
		di: collection,
	}
}
