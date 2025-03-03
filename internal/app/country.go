package app

import (
	"context"
	"fmt"
	"gateway_service/internal/domain"

	country_req "github.com/LiveisFpv/country_v1/gen/go/country"
)

func (a *App) Add_Country(ctx context.Context, countryTitle, countryCapital, countryArea string) (countryId int, err error) {
	req := &country_req.Add_Country_Request{
		CountryTitle:   countryTitle,
		CountryCapital: countryCapital,
		CountryArea:    countryArea,
	}

	resp, err := a.client_country.Add_Country(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("Add country failed: %w", err)
	}
	return int(resp.CountryId), nil
}

func (a *App) Get_CountryById(ctx context.Context, countryId int) (*domain.Country, error) {
	req := &country_req.Get_CountryById_Request{
		CountryId: int64(countryId),
	}

	resp, err := a.client_country.Get_CountryById(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Get country failed: %w", err)
	}
	return &domain.Country{
		Country_id:      int(resp.CountryId),
		Country_title:   resp.CountryTitle,
		Country_capital: resp.CountryCapital,
		Country_area:    resp.CountryArea,
	}, nil
}

func (a *App) Get_All_Country(ctx context.Context) (countries []*domain.Country, err error) {
	req := &country_req.Get_All_Country_Request{}

	resp, err := a.client_country.Get_All_Country(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Get all country failed: %w", err)
	}

	countries = []*domain.Country{}
	for _, country := range resp.Countries {
		countries = append(countries, &domain.Country{
			Country_id:      int(country.CountryId),
			Country_title:   country.CountryTitle,
			Country_capital: country.CountryCapital,
			Country_area:    country.CountryArea,
		})
	}
	return countries, nil
}

func (a *App) Update_CountryById(ctx context.Context, model *domain.Country) (country *domain.Country, err error) {
	req := &country_req.Update_CountryById_Request{
		CountryId:      int64(model.Country_id),
		CountryTitle:   model.Country_title,
		CountryCapital: model.Country_capital,
		CountryArea:    model.Country_area,
	}

	resp, err := a.client_country.Update_CountryById(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Update country failed: %w", err)
	}

	return &domain.Country{
		Country_id:      model.Country_id,
		Country_title:   resp.CountryTitle,
		Country_capital: resp.CountryCapital,
		Country_area:    resp.CountryArea,
	}, nil
}

func (a *App) Delete_CountryById(ctx context.Context, country_id int) (*domain.Country, error) {
	req := &country_req.Delete_CountryById_Request{
		CountryId: int64(country_id),
	}

	resp, err := a.client_country.Delete_CountryById(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Delete country failed: %w", err)
	}

	return &domain.Country{
		Country_title:   resp.CountryTitle,
		Country_capital: resp.CountryCapital,
		Country_area:    resp.CountryArea,
	}, nil
}
