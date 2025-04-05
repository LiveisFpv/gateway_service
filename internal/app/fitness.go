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
