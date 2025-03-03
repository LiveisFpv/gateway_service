package httpgin

import (
	"gateway_service/internal/app"
	"gateway_service/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Метод для получения страны (country)
// @Summary Получение страны
// @Description Возвращает страну
// @Tags country
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param country_id path int true "ID страны"
// @Success 200 {object} Get_CountryById_Response
// @Failure 400 {object}  Error_Response
// @Failure 401 {object}  Error_Response
// @Failure 500  {object}  Error_Response
// @Router /api/v1/country/{country_id} [get]
func getCountryById(c *gin.Context, a *app.App) {
	country_id, err := strconv.Atoi(c.Param("country_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	country, err := a.Get_CountryById(c, country_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, GetCountryByIdSuccessResponce(country))
}

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
// @Router /api/v2/country/all [post]
func getCountryAll(c *gin.Context, a *app.App) {
	var reqBody Get_All_Country_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	filter := []domain.Filter{}
	for _, filterRepr := range *reqBody.Filter {
		filter = append(filter, domain.Filter{
			Field: filterRepr.Field,
			Value: filterRepr.Value,
		})
	}

	sort := []domain.Sort{}
	for _, sortRepr := range *reqBody.Sort {
		sort = append(sort, domain.Sort{
			Direction: sortRepr.Direction,
			By:        sortRepr.By,
		})
	}

	queryResp, err := a.Get_All_Country(c,
		&domain.Pagination{
			Current: reqBody.Pagination.Current,
			Total:   reqBody.Pagination.Total,
			Limit:   reqBody.Pagination.Limit,
		},
		&filter,
		&sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, GetAllCountrySuccessResponse(queryResp, &reqBody.Pagination))
}

// Метод для добавления страны (country)
// @Summary Добавление страну
// @Description Добавляет страну
// @Tags country
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param country body Create_Country_Request true "Данные страны"
// @Success 200 {object} Create_Country_Response
// @Failure 400 {object}  Error_Response
// @Failure 401 {object}  Error_Response
// @Failure 500  {object}  Error_Response
// @Router /api/v1/country [post]
func createCountry(c *gin.Context, a *app.App) {
	var reqBody Create_Country_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	resp, err := a.Add_Country(c,
		reqBody.Data.Country_title,
		reqBody.Data.Country_capital,
		reqBody.Data.Country_area,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, CreateCountrySuccessResponce(resp))
}

// Метод для редактирования страны (country)
// @Summary Редактирование страны
// @Description Редактирует страну
// @Tags country
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param country body Update_CountryById_Request true "Данные страны"
// @Success 200 {object} Update_CountryById_Response
// @Failure 400 {object}  Error_Response
// @Failure 401 {object}  Error_Response
// @Failure 500  {object}  Error_Response
// @Router /api/v1/country [put]
func updateCountryById(c *gin.Context, a *app.App) {
	var reqBody Update_CountryById_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	resp, err := a.Update_CountryById(c, &domain.Country{
		Country_id:      reqBody.Data.Country_id,
		Country_title:   reqBody.Data.Country_title,
		Country_capital: reqBody.Data.Country_capital,
		Country_area:    reqBody.Data.Country_area,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, UpdateCountryByIdSuccessResponce(resp))
}

// Метод для удаления страны (country)
// @Summary Удаление страны
// @Description Удаляет страну
// @Tags country
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param country body Delete_CountryById_Request true "ID страны"
// @Success 200 {object} Delete_CountryById_Response
// @Failure 400 {object}  Error_Response
// @Failure 401 {object}  Error_Response
// @Failure 500  {object}  Error_Response
// @Router /api/v1/country [delete]
func deleteCountryById(c *gin.Context, a *app.App) {
	var reqBody Delete_CountryById_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	resp, err := a.Delete_CountryById(c, reqBody.Data.Country_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, DeleteCountryByIdSuccessResponce(resp))
}
