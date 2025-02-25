package httpgin

import (
	"gateway_service/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCountryById(c *gin.Context, a *app.App){
	var reqBody Get_CountryById_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	country, err := a.Get_CountryById(c, reqBody.Country_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, GetCountryByIdSuccessResponce(country))
}

func getCountryAll(c *gin.Context, a *app.App){
	countries, err := a.Get_All_Country(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, GetAllCountrySuccessResponse(countries))
}

func createCountry(c *gin.Context, a *app.App){
	var reqBody Create_Country_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	resp, err := a.Add_Country(c, reqBody.Country_title, reqBody.Country_capital, reqBody.Country_area)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, CreateCountryByIdSuccessResponce(resp))
}

func updateCountryById(c *gin.Context, a *app.App){
	var reqBody Update_CountryById_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	resp, err := a.Update_CountryById(c, reqBody.Country_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, UpdateCountryByIdSuccessResponce(resp))
}

func deleteCountryById(c *gin.Context, a *app.App){
	var reqBody Delete_CountryById_Request
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	resp, err := a.Delete_CountryById(c, reqBody.Country_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, DeleteCountryByIdSuccessResponce(resp))
}
