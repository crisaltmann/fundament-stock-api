package portfolio_service

import (
	"github.com/crisaltmann/fundament-stock-api/internal"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
)

type Service struct {
	repository Repository
}

type Repository interface {
	GetPortfolio(usuario int64) ([]portfolio_domain.Portfolio, error)
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

func (s Service) GetPortfolio(usuario int64) ([]portfolio_domain.Portfolio, error) {
	portfolio, err := s.repository.GetPortfolio(usuario)
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