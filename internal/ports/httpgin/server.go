package httpgin

import (
	"context"
	"gateway_service/internal/app"
	"gateway_service/internal/config"
	"gateway_service/internal/ports/httpgin/middlewares"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "gateway_service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	port       string
	app        *gin.Engine
	httpServer *http.Server
}

func NewHTTPServer(gincfg *config.GinConfig, a *app.App) *Server {
	//Use GIN how release
	gin.SetMode(gin.DebugMode)
	//TODO logger
	//Init clear gin server without logger and recovery
	r := gin.New()
	//Use custom logger
	r.Use(
		middlewares.RequestLogger(gincfg.Logger),
		middlewares.ResponseLogger(gincfg.Logger),
		gin.Recovery(),
	)
	httpServer := &http.Server{
		Addr:    gincfg.Port,
		Handler: r,
	}

	//Connect to server with default logger
	s := Server{port: gincfg.Port, app: r, httpServer: httpServer}
	//Enable all CORS policity if start on one machine
	s.app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Разрешаем запросы с любых доменов
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешаем все нужные методы
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	//Way to the open route
	open := s.app.Group("/")
	OpenRouter(open, a)
	//Way to protected route

	api := open.Group("api/v1")
	//Use Middleware to validate token
	api.Use(middlewares.AuthMiddleware())
	//For test
	apiv2 := open.Group("api/v2")
	apiv2.Use(middlewares.AuthMiddleware())
	TestRounter(apiv2, a)

	//TODO logger middleware?
	ProtectedRouter(api, a)

	s.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ! replace
	return &s
}

// Start GIN and it blocked if error happened
func (s *Server) Listen() error {
	//Can use tls after
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
