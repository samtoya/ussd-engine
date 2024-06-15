package contracts

type GatewayEngine interface {
	HandleRequest(interface{}) string
}

type GatewayFactory interface {
	Detect(string) GatewayEngine
}

type CacheRepository interface {
	GetByKey(string) (string, error)
	Connect()
	GetEngine() interface{}
}

type UssdScreen interface {
	RouteToScreen()
	Display() string
}

type UssdScreenBase struct {
}

func (b *UssdScreenBase) RouteToScreen() {

}

func (b *UssdScreenBase) Display() string {
	return ""
}
