package portfolio_service

import (
	portfolio_domain2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	portfolio_repository2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/repository"
)

type Service struct {
	Repository *portfolio_repository2.Repository
}

func (s Service) GetPortfolio(usuario string) ([]portfolio_domain2.Portfolio, error) {
	portfolio, err := s.Repository.GetPortfolio(usuario)
	if err != nil {
		return portfolio, err
	}
	return portfolio, nil
}