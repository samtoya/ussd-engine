package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"ussd-gateway/pkg/config"
	"ussd-gateway/pkg/controllers"
	"ussd-gateway/pkg/di"
	"ussd-gateway/pkg/factories"
	"ussd-gateway/pkg/gateways"
	"ussd-gateway/pkg/parsers"
	"ussd-gateway/pkg/repositories"
	"ussd-gateway/pkg/routers"
	"ussd-gateway/pkg/utils"
)

var (
	app      = echo.New()
	validate = validator.New()
	deps     = di.NewServiceCollection()
	parser   = parsers.NewYamlScreenParser()
	cfg      = config.New()
	rdb      = repositories.NewRedisClient(cfg)
)

func init() {
	cfg.LoadEnv()
	rdb.Connect()
	screens, err := parser.ParseDefault()
	if err != nil {
		log.Println("Failed to parse screen:", err)
		os.Exit(3)
	}
}

func main() {
	redisEngine := rdb.GetEngine().(*redis.Client)

	africasTalkingGateway := gateways.NewAfricasTalkingGateway(deps)
	arkesselGateway := gateways.NewArkesselGateway()
	cacheRepository := repositories.NewRedisCacheRepository(redisEngine)
	gatewayController := controllers.NewGatewayController(deps)
	gatewayFactory := factories.New(deps)
	router := routers.NewEngine(app)

	_ = deps.Register(utils.AfricasTalkingGateway, africasTalkingGateway)
	_ = deps.Register(utils.ArkesselGateway, arkesselGateway)
	_ = deps.Register(utils.CacheRepository, cacheRepository)
	_ = deps.Register(utils.GatewayController, gatewayController)
	_ = deps.Register(utils.GatewayFactory, gatewayFactory)
	_ = deps.Register(utils.Validator, validate)

	router.ApiRoutes(deps)

	app.Logger.Fatal(app.Start(cfg.GetAddr()))
}
