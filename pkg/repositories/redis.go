package repositories

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"ussd-gateway/pkg/config"
	"ussd-gateway/pkg/contracts"
)

type redisCacheRepository struct {
	engine *redis.Client
	cfg    config.AppConfig
	ctx    context.Context
}

func (repository *redisCacheRepository) GetByKey(k string) (string, error) {
	return repository.engine.Get(repository.ctx, k).Result()
}

func (repository *redisCacheRepository) Connect() {
	cacheOpts := repository.cfg.GetCache().(*redis.Options)
	repository.engine = redis.NewClient(cacheOpts)
}

func (repository *redisCacheRepository) GetEngine() interface{} {
	return repository.engine
}

func NewRedisClient(config config.AppConfig) contracts.CacheRepository {
	return &redisCacheRepository{
		cfg: config,
		ctx: context.Background(),
	}
}

func NewRedisCacheRepository(engine *redis.Client) contracts.CacheRepository {
	return &redisCacheRepository{engine: engine}
}
