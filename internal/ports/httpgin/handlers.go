package httpgin

import (
	"fmt"
	"gateway_service/internal/app"
	"gateway_service/internal/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Метод для получения картинок?
// @Summary Возвращает картинки
// @Description Возвращает картинки
// @Tags country
// @Accept json
// @Produce json
// @Success 200 {object} Get_Pictures_Response
// @Router /pictures [get]
func getPictures(c *gin.Context, a *app.App) {
	c.JSON(http.StatusOK, Get_Pictures_Response{
		&[]Pictures_ID_Response{
			{
				Title:       "Yoga",
				Description: "Yoga",
				ImageSrc:    "https://www.everydayyoga.com/cdn/shop/articles/yoga_1024x1024.jpg?v=1703853908",
			},
			{
				Title:       "Body pump",
				Description: "Body pump",
				ImageSrc:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQZS2jTw7KieqO57Xmu8YrgiaYp8gPxnD-tAQ&s",
			},
		},
	})
}

func getUser(c *gin.Context, a *app.App) {
	uidVal, exists := c.Keys["uid"]
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse(fmt.Errorf("uid not found")))
		return
	}

	user_id, ok := uidVal.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse(fmt.Errorf("uid is not an integer")))
		return
	}
	user, err := a.GetUser(c, int(user_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, UserSuccessResponse(user))
}

func createUser(c *gin.Context, a *app.App) {
	uidVal, exists := c.Keys["uid"]
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse(fmt.Errorf("uid not found")))
		return
	}

	user_id, ok := uidVal.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse(fmt.Errorf("uid is not an integer")))
		return
	}
	var reqBody UserPut
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	user, err := a.CreateUser(c,
		int(user_id),
		reqBody.User_firstName,
		reqBody.User_lastName,
		reqBody.User_middleName,
		reqBody.User_birthday,
		reqBody.User_height,
		reqBody.User_weight,
		reqBody.User_fitness_target,
		reqBody.User_sex,
		reqBody.User_level,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, UserSuccessResponse(user))
}

func updateUser(c *gin.Context, a *app.App) {
	uidVal, exists := c.Keys["uid"]
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse(fmt.Errorf("uid not found")))
		return
	}

	user_id, ok := uidVal.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse(fmt.Errorf("uid is not an integer")))
		return
	}
	var reqBody UserPut
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	user, err := a.UpdateUser(c,
		int(user_id),
		reqBody.User_firstName,
		reqBody.User_lastName,
		reqBody.User_middleName,
		reqBody.User_birthday,
		reqBody.User_height,
		reqBody.User_weight,
		reqBody.User_fitness_target,
		reqBody.User_sex,
		reqBody.User_level,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, UserSuccessResponse(user))
}

func GetPlanDiet(c *gin.Context, a *app.App) {
	c.JSON(http.StatusOK, Plan_diet_Response{
		Data: &[]Plan_diet{
			Plan_diet{
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
				Weight:      100,
				User_id:     1,
				Date:        (time.Now()).String(),
			},
			Plan_diet{
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
				Weight:      100,
				User_id:     1,
				Date:        (time.Now().Add(time.Hour * 24)).String(),
			},
			Plan_diet{
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
				Weight:      100,
				User_id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 2)).String(),
			},
			Plan_diet{
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
				Weight:      100,
				User_id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 3)).String(),
			},
			Plan_diet{
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
				Weight:      100,
				User_id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 4)).String(),
			},
			Plan_diet{
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
				Weight:      100,
				User_id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 5)).String(),
			},
			Plan_diet{
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
				Weight:      100,
				User_id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 6)).String(),
			},
		},
	})
}

func GetPlanTrain(c *gin.Context, a *app.App) {
	c.JSON(http.StatusOK, Plan_train_Response{
		Data: &[]Plan_train{
			Plan_train{
				Title:       "GYM",
				Description: "БЕГ",
				User_level:  1,
				User_Id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 0)).String(),
			},
			Plan_train{
				Title:       "GYM",
				Description: "ПРЕСС КАЧАТЬ",
				User_level:  1,
				User_Id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 1)).String(),
			},
			Plan_train{
				Title:       "GYM",
				Description: "ОТЖИМАНИЯ",
				User_level:  1,
				User_Id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 2)).String(),
			},
			Plan_train{
				Title:       "GYM",
				Description: "БЕГ",
				User_level:  1,
				User_Id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 3)).String(),
			},
			Plan_train{
				Title:       "GYM",
				Description: "ПРЕСС КАЧАТЬ",
				User_level:  1,
				User_Id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 4)).String(),
			},
			Plan_train{
				Title:       "GYM",
				Description: "ОТЖИМАНИЯ",
				User_level:  1,
				User_Id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 5)).String(),
			},
			Plan_train{
				Title:       "GYM",
				Description: "БЕГ",
				User_level:  1,
				User_Id:     1,
				Date:        (time.Now().Add(time.Hour * 24 * 6)).String(),
			},
		},
	})
}

func GetPlanHistory(c *gin.Context, a *app.App) {
	c.JSON(http.StatusOK, History_Response{
		&[]History{
			History{
				User_id: 1,
				Weight:  100,
				Date:    (time.Now().Add(-time.Hour * 24 * 6)).String(),
			},
			History{
				User_id: 1,
				Weight:  99,
				Date:    (time.Now().Add(-time.Hour * 24 * 3)).String(),
			},
			History{
				User_id: 1,
				Weight:  98,
				Date:    (time.Now().Add(-time.Hour * 24 * 1)).String(),
			},
		},
	})
}

func PutPlanHistory(c *gin.Context, a *app.App) {
	c.JSON(http.StatusOK, History{
		User_id: 1,
		Weight:  98,
		Date:    (time.Now()).String(),
	})
}

func GetDishesbyID(c *gin.Context, a *app.App) {
	id, err := strconv.Atoi(c.Param("dish_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	if id == 2 {
		c.JSON(http.StatusOK, Dishes_Response{
			Data: Dishes{
				Id:          2,
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
			},
		})
	} else if id == 3 {
		c.JSON(http.StatusOK, Dishes_Response{
			Data: Dishes{
				Id:          3,
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
			},
		})
	} else {
		c.JSON(http.StatusOK, Dishes_Response{
			Data: Dishes{
				Id:          1,
				Time:        "Завтрак",
				Title:       "Каша",
				Callory:     200,
				Fat:         5,
				Protein:     10,
				Carbs:       5,
				Description: "Вкусно",
			},
		})
	}

}

func GetTrainbyID(c *gin.Context, a *app.App) {
	id, err := strconv.Atoi(c.Param("train_id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	if id == 2 {
		c.JSON(http.StatusOK, Train_Response{
			Train{
				Id:          2,
				Title:       "GYM",
				Description: "ПРЕСС КАЧАТЬ",
				User_level:  1,
			},
		})
	} else if id == 3 {
		c.JSON(http.StatusOK, Train_Response{
			Train{
				Id:          3,
				Title:       "GYM",
				Description: "ОТЖИМАНИЯ",
				User_level:  1,
			},
		})
	} else {
		c.JSON(http.StatusOK, Train_Response{
			Train{
				Id:          1,
				Title:       "GYM",
				Description: "БЕГ",
				User_level:  1,
			},
		})
	}
}

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
// @Router /api/v1/country/all [post]
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
