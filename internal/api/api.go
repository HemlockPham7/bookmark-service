package api

import (
	"fmt"
	"net/http"

	"github.com/HemlockPham7/bookmark-service/docs"
	_ "github.com/HemlockPham7/bookmark-service/docs"
	bookmarkHdl "github.com/HemlockPham7/bookmark-service/internal/app/handler/bookmark"
	genCodeHdl "github.com/HemlockPham7/bookmark-service/internal/app/handler/gencode"
	healthcheckHdl "github.com/HemlockPham7/bookmark-service/internal/app/handler/healthcheck"
	linkHdl "github.com/HemlockPham7/bookmark-service/internal/app/handler/link"
	bookmarkRepo "github.com/HemlockPham7/bookmark-service/internal/app/repository/bookmark"
	"github.com/HemlockPham7/bookmark-service/internal/app/repository/cache"
	healthCheckRepo "github.com/HemlockPham7/bookmark-service/internal/app/repository/healthcheck"
	linkRepo "github.com/HemlockPham7/bookmark-service/internal/app/repository/link"
	mqRepo "github.com/HemlockPham7/bookmark-service/internal/app/repository/queue"
	bookmarkSvc "github.com/HemlockPham7/bookmark-service/internal/app/service/bookmark"
	healthCheckSvc "github.com/HemlockPham7/bookmark-service/internal/app/service/healthcheck"
	linkSvc "github.com/HemlockPham7/bookmark-service/internal/app/service/link"
	mqSvc "github.com/HemlockPham7/bookmark-service/internal/app/service/queue"
	"github.com/HemlockPham7/common-libs/pkg/jwtutils"
	"github.com/HemlockPham7/common-libs/pkg/middleware"
	"github.com/HemlockPham7/common-libs/pkg/ratelimitutils"
	"github.com/HemlockPham7/common-libs/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// Engine interface for starting the application
type Engine interface {
	Start() error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// engine struct for starting the application
type engine struct {
	app         *gin.Engine
	cfg         *Config
	redisClient *redis.Client
	dbClient    *gorm.DB
	jwtGen      jwtutils.JWTGenerator
	jwtVal      jwtutils.JWTValidator
}

type EngineOpts struct {
	App         *gin.Engine
	Cfg         *Config
	RedisClient *redis.Client
	DbClient    *gorm.DB
	JwtGen      jwtutils.JWTGenerator
	JwtVal      jwtutils.JWTValidator
}

// NewEngine creates a new engine
func NewEngine(opts *EngineOpts) Engine {
	app := &engine{
		app:         opts.App,
		cfg:         opts.Cfg,
		redisClient: opts.RedisClient,
		dbClient:    opts.DbClient,
		jwtGen:      opts.JwtGen,
		jwtVal:      opts.JwtVal,
	}
	app.initRoutes()
	return app
}

// Start starts the application
func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

// ServeHTTP to test the API endpoint
func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

type handlers struct {
	genCodeHandler     genCodeHdl.Handler
	linkHandler        linkHdl.Handler
	healthCheckHandler healthcheckHdl.Handler
	bookmarkHandler    bookmarkHdl.Handler
}

func (e *engine) initHandlers() *handlers {
	genCodeService := utils.NewGenCode()
	genCodeHandler := genCodeHdl.NewHandler(genCodeService)

	healthCheckRepository := healthCheckRepo.NewHealthCheckRepository(e.redisClient)
	healthCheckService := healthCheckSvc.NewHealthCheckService(e.cfg.ServiceName, e.cfg.InstanceID, healthCheckRepository)
	healthCheckHandler := healthcheckHdl.NewHealthcheckHandler(healthCheckService)

	// init cache db
	distributedCache := cache.NewRedisDB(e.redisClient)

	// init message queue
	mqRepository := mqRepo.NewRedisQueue(e.redisClient, e.cfg.QueueName)
	mqService := mqSvc.NewService(mqRepository)

	bookmarkRepository := bookmarkRepo.NewRepository(e.dbClient)
	bookmarkService := bookmarkSvc.NewService(bookmarkRepository, genCodeService)
	bookmarkServiceWithCache := bookmarkSvc.NewBookmarkServiceWithCache(bookmarkService, distributedCache)
	bookmarkHandler := bookmarkHdl.NewHandler(bookmarkServiceWithCache, mqService)

	// init link Service
	linkRepository := linkRepo.NewLinkRepository(e.redisClient)
	linkService := linkSvc.NewLinkService(linkRepository, bookmarkRepository, genCodeService)
	linkHandler := linkHdl.NewLinkHandler(linkService)

	return &handlers{
		genCodeHandler:     genCodeHandler,
		linkHandler:        linkHandler,
		healthCheckHandler: healthCheckHandler,
		bookmarkHandler:    bookmarkHandler,
	}
}

type middlewares struct {
	jwtAuth   middleware.JWTAuth
	rateLimit middleware.RateLimit
}

// initMiddlewares initializes the middlewares
func (e *engine) initMiddlewares() middlewares {
	jwtAuth := middleware.NewJWTAuth(e.jwtVal)

	rateLimitRepository := ratelimitutils.NewRedisRepo(e.redisClient)
	rateLimit := middleware.NewRateLimit(rateLimitRepository)

	return middlewares{
		jwtAuth:   jwtAuth,
		rateLimit: rateLimit,
	}
}

// initRoutes initializes the routes
func (e *engine) initRoutes() {
	allHandlers := e.initHandlers()
	allMiddlewares := e.initMiddlewares()

	// gencode
	e.app.GET("/gencode", allHandlers.genCodeHandler.GenerateCode)

	// health-check
	e.app.GET("/health-check", allHandlers.healthCheckHandler.HealthCheck)

	// Init swagger routes
	docs.SwaggerInfo.BasePath = e.cfg.BasePath
	e.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	privateRoutes := e.app.Group("")
	privateRoutes.Use(allMiddlewares.jwtAuth.JWTAuth())
	privateRoutes.Use(allMiddlewares.rateLimit.RateLimit())
	{
		privateV1Routes := privateRoutes.Group("/v1")
		{
			bookmarksRoutes := privateV1Routes.Group("/bookmarks")
			{
				bookmarksRoutes.POST("/", allHandlers.bookmarkHandler.CreateBookmark)
				bookmarksRoutes.GET("/", allHandlers.bookmarkHandler.GetBookmarks)
				bookmarksRoutes.POST("/import", allHandlers.bookmarkHandler.ImportBookmarks)
				bookmarksRoutes.DELETE("/:id", allHandlers.bookmarkHandler.DeleteBookmarkByID)
				bookmarksRoutes.PUT("/:id", allHandlers.bookmarkHandler.UpdateBookmarkByID)
			}
		}
	}

	publicRoutes := e.app.Group("")
	{
		publicV1Routes := publicRoutes.Group("/v1")
		{
			linksRoutes := publicV1Routes.Group("/links")
			{
				linksRoutes.POST("/shorten", allHandlers.linkHandler.ShortenLink)
				linksRoutes.GET("/redirect/:code", allHandlers.linkHandler.Redirect)
			}
		}
	}

}
