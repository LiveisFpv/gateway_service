package httpgin

import (
	"gateway_service/internal/app"
	"gateway_service/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Func send 401 if incorrect user password or email]
send 200 and token if it correct
*/
// Метод для авторизации
// @Summary Авторизация
// @Description Авторизует пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param auth body AuthRequest true "Данные пользователя"
// @Success 200 {object} AuthResponse
// @Failure 401 {object}  Error_Response
// @Failure 500  {object}  Error_Response
// @Router /auth [post]
func auth(c *gin.Context, a *app.App) {
	var reqBody AuthRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	//Its only way to the sun (auth-service)
	token, err := a.Auth(c, reqBody.Email, reqBody.Password, reqBody.YandexToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, AuthSuccessResponce(&domain.Token{
		Token: token,
	}))
}

/*
Func send 200 if user registered
send 500 if is problem on Auth-service
*/
// Метод для регистрации пользователя
// @Summary Регистрирует пользователя
// @Description Регистрирует пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RegisterRequest true "Данные пользователя"
// @Success 200 {} ok
// @Failure 500  {object}  Error_Response
// @Router /register [post]
func register(c *gin.Context, a *app.App) {
	var reqBody RegisterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	//Its only way to the sun (auth-service)
	err := a.Register(c, reqBody.Email, reqBody.Password, reqBody.YandexToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{
		Data: true,
	})
}
