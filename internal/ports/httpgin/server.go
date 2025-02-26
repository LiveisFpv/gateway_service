package httpgin

import (
	"gateway_service/internal/app"
	"gateway_service/internal/crypt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "gateway_service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	port string
	app  *gin.Engine
}

func NewHTTPServer(port string, a *app.App) Server {
	//Use GIN how release
	gin.SetMode(gin.ReleaseMode)
	//TODO logger
	//Connect to server with default logger
	s := Server{port: port, app: gin.Default()}
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
	api.Use(crypt.AuthMiddleware())
	//TODO logger middleware?
	ProtectedRouter(api, a)
	s.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // ! replace
	return s
}

// Start GIN and it blocked if error happened
func (s *Server) Listen() error {
	return s.app.Run(s.port)
}

// Only for tests
func (s *Server) Handler() http.Handler {
	return s.app
}
