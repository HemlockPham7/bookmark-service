package infrastructure

import (
	"github.com/HemlockPham7/bookmark-service/internal/api"
	"github.com/HemlockPham7/common-libs/pkg/logger"
	"github.com/gin-gonic/gin"
)

func CreateAPIConfig() *api.Config {
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

func CreateAPI() api.Engine {
	// create app config
	cfg := CreateAPIConfig()

	// set log level
	logger.SetLogLevel(cfg.LogLevel)

	// create redis client
	redisClient := CreateRedisClient("bookmark")

	// Init db
	db := CreateDB("bookmark")

	jwtGen, jwtVal := CreateJWTProvider()

	// create newrelic client
	nrClient := CreateNRClient()

	app := gin.Default()

	return api.NewEngine(&api.EngineOpts{
		App:         app,
		Cfg:         cfg,
		RedisClient: redisClient,
		DbClient:    db,
		JwtGen:      jwtGen,
		JwtVal:      jwtVal,
		NrClient:    nrClient,
	})
}
