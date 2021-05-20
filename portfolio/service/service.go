package portfolio_service

import (
	portfolio_domain "github.com/crisaltmann/fundament-stock-api/portfolio/domain"
	portfolio_repository "github.com/crisaltmann/fundament-stock-api/portfolio/repository"
)

type Service struct {
	Repository *portfolio_repository.Repository
}

func (s Service) GetPortfolio(usuario string) ([]portfolio_domain.Portfolio, error) {
	portfolio, err := s.Repository.GetPortfolio(usuario)
	if err != nil {
		return portfolio, err
	}
	return portfolio, nil
}