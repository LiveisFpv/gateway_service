package app

import (
	"context"
	"gateway_service/internal/domain"

	fitness_v1 "github.com/LiveisFPV/fitness_v1/gen/go/fitness"
)

func (a *App) GetUser(ctx context.Context, user_id int) (*domain.User, error) {
	req := &fitness_v1.ProfileRequest{
		UserId: int64(user_id),
	}
	resp, err := a.client_fitness.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		User_id:             user_id,
		User_firstName:      resp.UserFirstname,
		User_lastName:       resp.UserLastname,
		User_middleName:     resp.UserMiddlename,
		User_birthday:       resp.UserBirthday,
		User_height:         int(resp.UserHeight),
		User_weight:         resp.UserWeight,
		User_fitness_target: resp.UserFitnessTarget,
		User_sex:            resp.UserSex,
		User_level:          int(resp.UserLevel),
	}
	return user, nil
}

func (a *App) CreateUser(
	ctx context.Context,
	user_id int,
	user_firstName *string,
	user_lastName *string,
	user_middleName *string,
	user_birthday *string,
	user_height *int32,
	user_weight *float64,
	user_fitness_target *string,
	user_sex *bool,
	user_level *int32,
) (*domain.User, error) {
	req := &fitness_v1.CreateProfileRequest{
		UserId:            int64(user_id),
		UserFirstname:     *user_firstName,
		UserLastname:      *user_lastName,
		UserMiddlename:    user_middleName,
		UserBirthday:      *user_birthday,
		UserHeight:        *user_height,
		UserWeight:        *user_weight,
		UserFitnessTarget: *user_fitness_target,
		UserSex:           *user_sex,
		UserLevel:         user_level,
	}
	resp, err := a.client_fitness.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		User_id:             user_id,
		User_firstName:      resp.UserFirstname,
		User_lastName:       resp.UserLastname,
		User_middleName:     resp.UserMiddlename,
		User_birthday:       resp.UserBirthday,
		User_height:         int(resp.UserHeight),
		User_weight:         resp.UserWeight,
		User_fitness_target: resp.UserFitnessTarget,
		User_sex:            resp.UserSex,
		User_level:          int(resp.UserLevel),
	}
	return user, nil
}

func (a *App) UpdateUser(
	ctx context.Context,
	user_id int,
	user_firstName *string,
	user_lastName *string,
	user_middleName *string,
	user_birthday *string,
	user_height *int32,
	user_weight *float64,
	user_fitness_target *string,
	user_sex *bool,
	user_level *int32,
) (*domain.User, error) {
	req := &fitness_v1.UpdateProfileRequest{
		UserId:            int64(user_id),
		UserFirstname:     user_firstName,
		UserLastname:      user_lastName,
		UserMiddlename:    user_middleName,
		UserBirthday:      user_birthday,
		UserHeight:        user_height,
		UserWeight:        user_weight,
		UserFitnessTarget: user_fitness_target,
		UserSex:           user_sex,
		UserLevel:         user_level,
	}
	resp, err := a.client_fitness.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		User_id:             user_id,
		User_firstName:      resp.UserFirstname,
		User_lastName:       resp.UserLastname,
		User_middleName:     resp.UserMiddlename,
		User_birthday:       resp.UserBirthday,
		User_height:         int(resp.UserHeight),
		User_weight:         resp.UserWeight,
		User_fitness_target: resp.UserFitnessTarget,
		User_sex:            resp.UserSex,
		User_level:          int(resp.UserLevel),
	}
	return user, nil
}

func (a *App) GetPlanDishes(ctx context.Context, user_id int, date string) (*domain.Plan_diet, error) {
	req := &fitness_v1.GetPlanDishesRequest{
		UserId: int64(user_id),
		Date:   date,
	}
	resp, err := a.client_fitness.GetPlanDishes(ctx, req)
	if err != nil {
		return nil, err
	}
	plan_diet := &domain.Plan_diet{
		Data: &[]domain.Diet{},
	}
	for _, el := range resp.Data {
		*plan_diet.Data = append(*plan_diet.Data, domain.Diet{
			Dishes_id:   int(el.DishesId),
			Time:        el.Time,
			Title:       el.DishesTitle,
			Callory:     el.Kcal,
			Fat:         el.Fat,
			Protein:     el.Protein,
			Carbs:       el.Carbs,
			Description: el.Description,
			Weight:      el.DishesWeight,
			Date:        el.Date,
		})
	}
	return plan_diet, nil
}

func (a *App) GetPlanTrain(ctx context.Context, user_id int, date string) (*domain.Plan_train, error) {
	req := &fitness_v1.GetPlanTrainRequest{
		UserId: int64(user_id),
		Date:   date,
	}
	resp, err := a.client_fitness.GetPlanTrain(ctx, req)
	if err != nil {
		return nil, err
	}
	plan_train := &domain.Plan_train{
		Data: &[]domain.Train{},
	}
	for _, el := range resp.Data {
		*plan_train.Data = append(*plan_train.Data, domain.Train{
			Train_id:    int(el.TrainId),
			Title:       el.TrainTitle,
			Description: el.TrainDescription,
			User_level:  int(el.UserLevel),
			Date:        el.Date,
		})
	}
	return plan_train, nil
}

func (a *App) GetHistory(ctx context.Context, user_id int, date string) (*domain.WeightHistory, error) {
	req := &fitness_v1.GetHistoryRequest{
		UserId: int64(user_id),
		Date:   date,
	}
	resp, err := a.client_fitness.GetHistory(ctx, req)
	if err != nil {
		return nil, err
	}
	history := &domain.WeightHistory{
		Data: &[]domain.WeightRecord{},
	}
	for _, el := range resp.Data {
		*history.Data = append(*history.Data, domain.WeightRecord{
			User_id: user_id,
			Weight:  float32(el.UserWeight),
			Date:    el.Date,
		})
	}
	return history, nil
}

// func (a *App) PutHistory(ctx context.Context, user_id int, weight float32, date string) (*domain.WeightRecord, error) {
// 	req := &fitness_v1.History{
// 		UserId: int64(user_id),
// 		UserWeight: float64(weight),
// 		Date:   date,
// 	}
// 	resp, err := a.client_fitness.PutHistory(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	record := &domain.WeightRecord{
// 		User_id: user_id,
// 		Weight: resp.weight,
// 	}
// }
