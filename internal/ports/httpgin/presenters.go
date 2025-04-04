package httpgin

import (
	"gateway_service/internal/domain"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email       string  `json:"email"`
	Password    *string `json:"password"`
	YandexToken *string `json:"yandex_token"`
}

type UserDataRequest struct {
	UID           int64   `json:"id"`
	Birthday      string  `json:"birthday"`
	Height        int     `json:"height"`
	Weight        float64 `json:"weight"`
	FitnessTarget string  `json:"fitness_target"`
	Sex           bool    `json:"sex"`
	Hypertain     bool    `json:"hypertain"`
	Diabet        bool    `json:"diabet"`
}

type Get_Pictures_Response struct {
	Data *[]Pictures_ID_Response
}

type Pictures_ID_Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageSrc    string `json:"imageSrc"`
}

type AuthRequest struct {
	Email       string  `json:"email"`
	Password    *string `json:"password"`
	YandexToken *string `json:"yandex_token"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type Get_CountryById_Request struct {
	Data *CountryIDRepr `json:"data"`
}

type CountryRepr struct {
	Country_id      int    `json:"country_id"`
	Country_title   string `json:"country_title"`
	Country_capital string `json:"country_capital"`
	Country_area    string `json:"country_area"`
}

type CountryIDRepr struct {
	Country_id int `json:"country_id"`
}

type CountryDataRepr struct {
	Country_title   string `json:"country_title"`
	Country_capital string `json:"country_capital"`
	Country_area    string `json:"country_area"`
}

type Get_CountryById_Response struct {
	Data *CountryRepr `json:"data"`
}

type PaginationRepr struct {
	Current int `json:"current"`
	Total   int `json:"total"`
	Limit   int `json:"limit"`
}

type FilterRepr struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type SortRepr struct {
	Direction string `json:"direction"`
	By        string `json:"by"`
}

type Get_All_Country_Request struct {
	Filter     *[]FilterRepr  `json:"filter"`
	Pagination PaginationRepr `json:"pagination"`
	Sort       *[]SortRepr    `json:"sort"`
}

type Get_All_Country_Response struct {
	Data       *[]CountryRepr  `json:"data"`
	Pagination *PaginationRepr `json:"pagination"`
}

type Create_Country_Request struct {
	Data *CountryDataRepr `json:"data"`
}

type Create_Country_Response struct {
	Data *CountryIDRepr `json:"data"`
}

type Update_CountryById_Request struct {
	Data *CountryRepr `json:"data"`
}

type Update_CountryById_Response struct {
	Data *CountryDataRepr `json:"data"`
}

type Delete_CountryById_Request struct {
	Data *CountryIDRepr `json:"data"`
}

type Delete_CountryById_Response struct {
	Data *CountryDataRepr `json:"data"`
}

// Interface for all responses
type ResponseData interface{}

type Error_Response struct {
	Data  *interface{} `json:"data"`
	Error string       `json:"error"`
}

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

// Response parsers
func AuthSuccessResponce(token *domain.Token) *gin.H {
	response := AuthResponse{
		Token: token.Token,
	}
	return SuccessResponse(response)
}

func GetCountryByIdSuccessResponce(country *domain.Country) *Get_CountryById_Response {
	return &Get_CountryById_Response{
		Data: &CountryRepr{
			Country_id:      country.Country_id,
			Country_title:   country.Country_title,
			Country_capital: country.Country_capital,
			Country_area:    country.Country_area,
		},
	}
}

func UpdateCountryByIdSuccessResponce(country *domain.Country) *Update_CountryById_Response {
	return &Update_CountryById_Response{
		Data: &CountryDataRepr{
			Country_title:   country.Country_title,
			Country_capital: country.Country_capital,
			Country_area:    country.Country_area,
		},
	}
}
func CreateCountrySuccessResponce(Country_id int) *Create_Country_Response {
	return &Create_Country_Response{
		Data: &CountryIDRepr{
			Country_id: Country_id,
		},
	}
}

func DeleteCountryByIdSuccessResponce(country *domain.Country) *Delete_CountryById_Response {
	return &Delete_CountryById_Response{
		Data: &CountryDataRepr{
			Country_title:   country.Country_title,
			Country_capital: country.Country_capital,
			Country_area:    country.Country_area,
		},
	}
}

func GetAllCountrySuccessResponse(countries *[]domain.Country, pagination *PaginationRepr) *Get_All_Country_Response {
	// Parsing data
	dat := []CountryRepr{}
	for _, country := range *countries {
		dat = append(dat, CountryRepr{
			Country_id:      country.Country_id,
			Country_title:   country.Country_title,
			Country_capital: country.Country_capital,
			Country_area:    country.Country_area,
		})
	}

	return &Get_All_Country_Response{
		Data:       &dat,
		Pagination: pagination,
	}
}
