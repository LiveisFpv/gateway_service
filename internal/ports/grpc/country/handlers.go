package country

import (
	"context"

	"github.com/LiveisFpv/country_v1/gen/go/country"
)

func (c *Client) Add_Country(
	ctx context.Context,
	req *country.Add_Country_Request,
) (
	resp *country.Add_Country_Response,
	err error) {
	return c.api.Add_Country(ctx, req)
}

func (c *Client) Get_CountryById(ctx context.Context, req *country.Get_CountryById_Request) (resp *country.Get_CountryById_Response, err error) {
	return c.api.Get_CountryById(ctx, req)
}
func (c *Client) Get_All_Country(ctx context.Context, req *country.Get_All_Country_Request) (resp *country.Get_All_Country_Response, err error) {
	return c.api.Get_All_Country(ctx, req)
}
func (c *Client) Update_CountryById(ctx context.Context, req *country.Update_CountryById_Request) (resp *country.Update_CountryById_Response, err error) {
	return c.api.Update_CountryById(ctx, req)
}
func (c *Client) Delete_CountryById(ctx context.Context, req *country.Delete_CountryById_Request) (resp *country.Delete_CountryById_Response, err error) {
	return c.api.Delete_CountryById(ctx, req)
}
