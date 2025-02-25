package app

import (
	"context"
	"fmt"
	"gateway_service/internal/domain"
	"gateway_service/internal/ports/grpc/authorization"
	"gateway_service/internal/ports/grpc/country"
	"gateway_service/internal/repository"

	country_req "github.com/LiveisFpv/country_v1/gen/go/country"

	"github.com/LiveisFPV/sso_v1/gen/go/sso"
)

type App struct {
	repo        repository.Repository
	client_auth *authorization.Client
	client_country *country.Client
}

func NewApp(repo repository.Repository, client_auth *authorization.Client) *App {
	return &App{
		repo:        repo,
		client_auth: client_auth,
	}
}

// ? Кто ты воин
func (a *App) Auth(ctx context.Context, email, password string) (string, error) {
	req := &sso.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    1,
	}

	resp, err := a.client_auth.Login(ctx, req)
	if err != nil {
		return "", fmt.Errorf("login failed: %w", err)
	}

	return resp.Token, nil
}

func (a *App) Register(ctx context.Context, email, password string) error {
	req := &sso.RegisterRequest{
		Email:    email,
		Password: password,
	}

	_, err := a.client_auth.Register(ctx, req)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}
	return nil
}

// TODO: Куда-то эти фигулечки убрать
func (a *App) Add_Country(ctx context.Context, countryTitle, countryCapital, countryArea string) (countryId int, err error){
	req := &country_req.Add_Country_Request{
		CountryTitle: countryTitle,
		CountryCapital: countryCapital,
		CountryArea: countryArea,
	}

	resp, err := a.client_country.Add_Country(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("Add country failed: %w", err)
	}
	return int(resp.CountryId), nil 
}

func (a *App) Get_CountryById(ctx context.Context, countryId int)(*domain.Country, error){
	req := &country_req.Get_CountryById_Request{
		CountryId: int64(countryId),
	}

	resp, err := a.client_country.Get_CountryById(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Get country failed: %w", err)
	}
	return &domain.Country{
		Country_id: int(resp.CountryId),
		Country_title: resp.CountryTitle,
		Country_capital: resp.CountryCapital,
		Country_area: resp.CountryArea,
	}, nil
}

func (a *App) Get_All_Country(ctx context.Context)(countries []*domain.Country, err error){
	req := &country_req.Get_All_Country_Request{}

	resp, err := a.client_country.Get_All_Country(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Get all country failed: %w", err)
	}

	countries = []*domain.Country{}
	for _, country := range resp.Countries {
		countries = append(countries, &domain.Country{
			Country_id: int(country.CountryId),
			Country_title: country.CountryTitle,
			Country_capital: country.CountryCapital,
			Country_area: country.CountryArea,
		})
	}
	return countries, nil
}

func (a *App) Update_CountryById(ctx context.Context, country_id int)(country *domain.Country, err error){
	req := &country_req.Update_CountryById_Request{
		CountryId: int64(country_id),
	}

	resp, err := a.client_country.Update_CountryById(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Update country failed: %w", err)
	}

	return &domain.Country{
		Country_id: country_id,
		Country_title: resp.CountryTitle,
		Country_capital: resp.CountryCapital,
		Country_area: resp.CountryArea,
	}, nil
}

func (a *App) Delete_CountryById(ctx context.Context, country_id int)(country_title string, err error){
	req := &country_req.Delete_CountryById_Request{
		CountryId: int64(country_id),
	}

	resp, err := a.client_country.Delete_CountryById(ctx, req)
	if err != nil {
		return "", fmt.Errorf("Delete country failed: %w", err)
	}

	return resp.CountryTitle, nil
}
