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

type Get_CountryById_Request struct {
	Country_id int
}

type Get_CountryById_Response struct {
	Country_id      int
	Country_title   string
	Country_capital string
	Country_area    string
}
type Pagination struct {
	Current int
	Total   int
	Limit   int
}

type Filter struct {
	Text string
}

type Sort struct {
	Direction string
	By        string
}
type Get_All_Country_Request struct {
	Filter     Filter
	Pagination Pagination
	Sort       Sort
}

type Get_All_Country_Response struct {
	Data       *[]Get_CountryById_Response
	Pagination *Pagination
	Error      *string
}

type Create_Country_Request struct {
	Country_title   string
	Country_capital string
	Country_area    string
}

type Create_Country_Response struct {
	Country_id int
}

type Update_CountryById_Request struct {
	Country_id      int
	Country_title   string
	Country_capital string
	Country_area    string
}

type Update_CountryById_Response struct {
	Country_title   string
	Country_capital string
	Country_area    string
}

type Delete_CountryById_Request struct {
	Country_id int
}

type Delete_CountryById_Response struct {
	Country_title string
}

// Interface for all responses
type ResponseData interface{}

type Error_Response struct {
	Data  *interface{}
	Error string
}

func ErrorResponse(err error) *gin.H {
	return &gin.H{
		"data":  nil,
		"error": err.Error(),
	}
}

// ! TODO Translte all response to form:
// data: Response
// error: string
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

func GetCountryByIdSuccessResponce(country *domain.Country) *gin.H {
	response := Get_CountryById_Response{
		Country_id:      country.Country_id,
		Country_title:   country.Country_title,
		Country_capital: country.Country_capital,
		Country_area:    country.Country_area,
	}
	return SuccessResponse(response)
}

func UpdateCountryByIdSuccessResponce(country *domain.Country) *gin.H {
	response := Update_CountryById_Response{
		Country_title:   country.Country_title,
		Country_capital: country.Country_capital,
		Country_area:    country.Country_area,
	}
	return SuccessResponse(response)
}
func CreateCountryByIdSuccessResponce(Country_id int) *gin.H {
	response := Create_Country_Response{
		Country_id: Country_id,
	}
	return SuccessResponse(response)
}
func DeleteCountryByIdSuccessResponce(Country_title string) *gin.H {
	response := Delete_CountryById_Response{
		Country_title: Country_title,
	}
	return SuccessResponse(response)
}
func GetAllCountrySuccessResponse(countries []*domain.Country) *gin.H {
	var response []Get_CountryById_Response
	for _, country := range countries {
		response = append(response, Get_CountryById_Response{
			Country_id:      country.Country_id,
			Country_title:   country.Country_title,
			Country_capital: country.Country_capital,
			Country_area:    country.Country_area,
		})
	}
	return SuccessResponse(response)
}
