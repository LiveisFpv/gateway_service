package fitness

import (
	"context"

	fitness "github.com/LiveisFPV/fitness_v1/gen/go/fitness"
)

func (c *Client) GetUser(
	ctx context.Context,
	req *fitness.ProfileRequest,
) (
	resp *fitness.ProfileResponse,
	err error,
) {
	return c.api.GetUser(ctx, req)
}

func (c *Client) CreateUser(
	ctx context.Context,
	req *fitness.CreateProfileRequest,
) (
	resp *fitness.ProfileResponse,
	err error,
) {
	return c.api.CreateUser(ctx, req)
}

func (c *Client) UpdateUser(
	ctx context.Context,
	req *fitness.UpdateProfileRequest,
) (
	resp *fitness.ProfileResponse,
	err error,
) {
	return c.api.UpdateUser(ctx, req)
}
func (c *Client) GetPlanDishes(ctx context.Context, in *fitness.GetPlanDishesRequest) (*fitness.PlanDishesResponse, error) {
	return c.api.GetPlanDishes(ctx, in)
}
func (c *Client) GetPlanTrain(ctx context.Context, in *fitness.GetPlanTrainRequest) (*fitness.PlanTrainResponse, error) {
	return c.api.GetPlanTrain(ctx, in)
}
func (c *Client) GetHistory(ctx context.Context, in *fitness.GetHistoryRequest) (*fitness.HistoryResponse, error) {
	return c.api.GetHistory(ctx, in)
}
func (c *Client) GetTrainInstr(ctx context.Context, in *fitness.GetTrainInstrRequest) (*fitness.TrainInstrResponse, error) {
	return c.api.GetTrainInstr(ctx, in)
}
func (c *Client) GetRecipe(ctx context.Context, in *fitness.GetRecipeRequest) (*fitness.RecipeResponse, error) {
	return c.api.GetRecipe(ctx, in)
}
