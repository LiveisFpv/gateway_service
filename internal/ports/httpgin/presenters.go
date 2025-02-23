package httpgin

import (
	"gateway_service/internal/domain"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

// Interface for all responses
type ResponseData interface{}

func ErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

func SuccessResponse(data ResponseData) *gin.H {
	return &gin.H{
		"data":  data,
		"error": nil,
	}
}

func AuthSuccessResponce(token *domain.Token) *gin.H {
	response := AuthResponse{
		Token: token.Token,
	}
	return SuccessResponse(response)
}
