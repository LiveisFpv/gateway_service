package httpgin

import (
	"gateway_service/internal/app"

	"github.com/gin-gonic/gin"
)

// Use AuthMiddleware before progress request
func ProtectedRouter(r *gin.RouterGroup, a *app.App) {

}

// Open Router that getted a auth for user, without other functional API
func OpenRouter(r *gin.RouterGroup, a *app.App) {
	r.POST("/register", func(c *gin.Context) { register(c, a) })
	r.POST("/auth", func(c *gin.Context) { auth(c, a) })
}
