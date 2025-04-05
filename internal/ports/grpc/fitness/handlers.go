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
