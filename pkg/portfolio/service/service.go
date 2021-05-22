package portfolio_service

import (
	"github.com/crisaltmann/fundament-stock-api/internal"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
)

type Service struct {
	Repository Repository
}

type Repository interface {
	GetPortfolio(usuario string) ([]portfolio_domain.Portfolio, error)
}

func NewService(repository Repository) Service {
	return Service{Repository: repository}
}

func (s Service) GetPortfolio(usuario string) ([]portfolio_domain.Portfolio, error) {
	portfolio, err := s.Repository.GetPortfolio(usuario)
	if err != nil {
		return portfolio, err
	}
	for idx, item := range portfolio {
		if item.Ativo.Cotacao > 0 {
			portfolio[idx].Valor = internal.RoundFloat(float32(item.Quantidade) * item.Ativo.Cotacao)
		}
	}
	return portfolio, nil
}