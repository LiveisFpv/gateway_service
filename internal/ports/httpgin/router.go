package httpgin

import (
	"gateway_service/internal/app"

	"github.com/gin-gonic/gin"
)

// Use AuthMiddleware before progress request
func ProtectedRouter(r *gin.RouterGroup, a *app.App) {
	r.GET("/country/:country_id", func(c *gin.Context) { getCountryById(c, a) })
	r.POST("/country/all", func(c *gin.Context) { getCountryAll(c, a) })
	r.POST("/country", func(c *gin.Context) { createCountry(c, a) })
	r.PUT("/country", func(c *gin.Context) { updateCountryById(c, a) })
	r.DELETE("/country", func(c *gin.Context) { deleteCountryById(c, a) })

	r.GET("/user", func(c *gin.Context) { getUser(c, a) })
	r.POST("/edit-profile", func(c *gin.Context) { createUser(c, a) })
	r.PUT("/edit-profile", func(c *gin.Context) { updateUser(c, a) })

	r.GET("/diet-plan", func(c *gin.Context) { GetPlanDiet(c, a) })
	r.GET("/training-plan", func(c *gin.Context) { GetPlanTrain(c, a) })
	r.GET("/history", func(c *gin.Context) { GetPlanHistory(c, a) })
	// r.PUT("/user/history", func(c *gin.Context) { PutPlanHistory(c, a) })
	r.GET("/dishes/cooking/:dish_id", func(c *gin.Context) { GetDishesbyID(c, a) })
	r.GET("/train/instruction/:train_id", func(c *gin.Context) { GetTrainbyID(c, a) })

}
func TestRounter(r *gin.RouterGroup, a *app.App) {
	// r.POST("/country", func(c *gin.Context) { new_getCountryAll(c, a) })
}

// Open Router that getted a auth for user, without other functional API
func OpenRouter(r *gin.RouterGroup, a *app.App) {
	r.POST("/register", func(c *gin.Context) { register(c, a) })
	r.POST("/auth", func(c *gin.Context) { auth(c, a) })
	r.GET("/country", func(c *gin.Context) { getCountryAll(c, a) })
	r.GET("/pictures", func(c *gin.Context) { getPictures(c, a) })
}
