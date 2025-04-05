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
