package httpgin

import (
	"gateway_service/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Метод для получения стран (country)
// @Summary Получение стран
// @Description Возвращает страны
// @Tags country
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param data body Get_All_Country_Request true "Filter,Pagination,Sort"
// @Success 200 {object} Get_All_Country_Response
// @Failure 401 {object}  Error_Response
// @Failure 500  {object}  Error_Response
// @Router /api/v2/country [post]
func new_getCountryAll(c *gin.Context, a *app.App) {
	countries, err := a.Get_All_Country(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	//Заглушка до реализации
	var response []Get_CountryById_Response
	for _, country := range countries {
		response = append(response, Get_CountryById_Response{
			Country_id:      country.Country_id,
			Country_title:   country.Country_title,
			Country_capital: country.Country_capital,
			Country_area:    country.Country_area,
		})
	}

	c.JSON(http.StatusOK, &Get_All_Country_Response{
		Data: &response,
		Pagination: &Pagination{
			Current: 2,
			Total:   20,
			Limit:   5,
		},
		Error: nil,
	})
}
