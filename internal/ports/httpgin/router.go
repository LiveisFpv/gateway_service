package httpgin

import (
	"gateway_service/internal/app"

	"github.com/gin-gonic/gin"
)

// Use AuthMiddleware before progress request
func ProtectedRouter(r *gin.RouterGroup, a *app.App) {
	r.GET("/country/:country_id", func(c *gin.Context) { getCountryById(c, a) })
	r.GET("/country", func(c *gin.Context) { getCountryAll(c, a) })
	r.POST("/country", func(c *gin.Context) { createCountry(c, a) })
	r.PUT("/country", func(c *gin.Context) { updateCountryById(c, a) })
	r.DELETE("/country", func(c *gin.Context) { deleteCountryById(c, a) })
}
func TestRounter(r *gin.RouterGroup, a *app.App) {
	r.GET("/country", func(c *gin.Context) { new_getCountryAll(c, a) })
}

// Open Router that getted a auth for user, without other functional API
func OpenRouter(r *gin.RouterGroup, a *app.App) {
	r.POST("/register", func(c *gin.Context) { register(c, a) })
	r.POST("/auth", func(c *gin.Context) { auth(c, a) })
}
