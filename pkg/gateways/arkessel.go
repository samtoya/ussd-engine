package gateways

import "ussd-gateway/pkg/contracts"

type arkesselGateway struct {
}

func (gtw arkesselGateway) HandleRequest(req interface{}) string {
	//TODO implement me
	panic("implement me")
}

func NewArkesselGateway() contracts.GatewayEngine {
	return &arkesselGateway{}
}
