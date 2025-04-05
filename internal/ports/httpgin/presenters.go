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

type UserPut struct {
	User_id             *int     `json:"user_id            "`
	User_firstName      *string  `json:"user_firstName     "`
	User_lastName       *string  `json:"user_lastName      "`
	User_middleName     *string  `json:"user_middleName    "`
	User_birthday       *string  `json:"user_birthday      "`
	User_height         *int32   `json:"user_height        "`
	User_weight         *float64 `json:"user_weight        "`
	User_fitness_target *string  `json:"user_fitness_target"`
	User_sex            *bool    `json:"user_sex           "`
	User_level          *int32   `json:"user_level         "`
}

type User struct {
	User_id             int     `json:"user_id            "`
	User_firstName      string  `json:"user_firstName     "`
	User_lastName       string  `json:"user_lastName      "`
	User_middleName     *string `json:"user_middleName    "`
	User_birthday       string  `json:"user_birthday      "`
	User_height         int     `json:"user_height        "`
	User_weight         float64 `json:"user_weight        "`
	User_fitness_target string  `json:"user_fitness_target"`
	User_sex            bool    `json:"user_sex           "`
	User_level          int     `json:"user_level         "`
}

type UserResponse struct {
	Data User
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

type Dishes struct {
	Id          int     `json:"id"`
	Time        string  `json:"time"`
	Title       string  `json:"title"`
	Callory     float64 `json:"callory"`
	Fat         float64 `json:"fat"`
	Protein     float64 `json:"protein"`
	Carbs       float64 `json:"carbs"`
	Description string  `json:"description"`
}

type Train struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User_level  int    `json:"user_level"`
}

type Plan_diet struct {
	Dishes_id int     `json:"dishes_id"`
	Weight    float64 `json:"weight"`
	User_id   int     `json:"user_id"`
	Date      string  `json:"date"`
}

type Plan_train struct {
	Train_id int    `json:"train_id"`
	User_Id  int    `json:"user_id"`
	Date     string `json:"date"`
}

type History struct {
	User_id int     `json:"user_id"`
	Weight  float64 `json:"weight"`
	Date    string  `json:"date"`
}

type Dishes_Response struct {
	Data Dishes `json:"data"`
}

type Train_Response struct {
	Data Train `json:"data"`
}

type Plan_diet_Response struct {
	Data *[]Plan_diet `json:"data"`
}

type Plan_train_Response struct {
	Data *[]Plan_train `json:"data"`
}

type History_Response struct {
	Data *[]History `json:"data"`
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
func UserSuccessResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		Data: User{
			User_id:             user.User_id,
			User_firstName:      user.User_firstName,
			User_lastName:       user.User_lastName,
			User_middleName:     user.User_middleName,
			User_birthday:       user.User_birthday,
			User_height:         user.User_height,
			User_weight:         user.User_weight,
			User_fitness_target: user.User_fitness_target,
			User_sex:            user.User_sex,
			User_level:          user.User_level,
		}}
}
