package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
)

type AppConfig interface {
	GetCache() interface{}
	LoadEnv()
	GetAddr() string
}

type appConfig struct {
}

func (cf *appConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", os.Getenv("APP.HOST"), os.Getenv("APP.PORT"))
}

func (cf *appConfig) LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	log.Println(".env variables loaded for environment:", os.Getenv("APP.ENV"))
}

func (cf *appConfig) GetCache() interface{} {
	db, _ := strconv.Atoi(os.Getenv("CACHE.DB"))
	cacheUrl := os.Getenv("CACHE.URL")
	if cacheUrl != "" {
		opts, err := redis.ParseURL(cacheUrl)
		if err != nil {
			log.Println("Cache: Failed to parse cache URL:", err)
			os.Exit(2)
		}

		return opts
	}

	opts := &redis.Options{
		Network:                    "",
		Addr:                       fmt.Sprintf("%s:%s", os.Getenv("CACHE.HOST"), os.Getenv("CACHE.PORT")),
		ClientName:                 os.Getenv("CACHE.CLIENT_NAME"),
		Dialer:                     nil,
		OnConnect:                  nil,
		Protocol:                   0,
		Username:                   os.Getenv("CACHE.USERNAME"),
		Password:                   os.Getenv("CACHE.PASSWORD"),
		CredentialsProvider:        nil,
		CredentialsProviderContext: nil,
		DB:                         db,
		MaxRetries:                 0,
		MinRetryBackoff:            0,
		MaxRetryBackoff:            0,
		DialTimeout:                0,
		ReadTimeout:                0,
		WriteTimeout:               0,
		ContextTimeoutEnabled:      false,
		PoolFIFO:                   false,
		PoolSize:                   0,
		PoolTimeout:                0,
		MinIdleConns:               0,
		MaxIdleConns:               0,
		MaxActiveConns:             0,
		ConnMaxIdleTime:            0,
		ConnMaxLifetime:            0,
		TLSConfig:                  nil,
		Limiter:                    nil,
		DisableIndentity:           false,
		IdentitySuffix:             "",
	}

	return opts
}

func New() AppConfig {
	return &appConfig{}
}
