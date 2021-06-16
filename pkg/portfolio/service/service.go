package portfolio_service

import (
	"github.com/crisaltmann/fundament-stock-api/internal"
	"github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
	"github.com/rs/zerolog/log"
	"time"
)

type Service struct {
	repository Repository
	quarterService QuarterService
}

type Repository interface {
	GetPortfolio(usuario int64, dataFinal time.Time) ([]portfolio_domain.Portfolio, error)
}

type QuarterService interface {
	GetQuarter(id int64) (quarter_domain.Trimestre, error)
}

func NewService(repository Repository, quarterService QuarterService) Service {
	return Service{
		repository: repository,
		quarterService: quarterService,
	}
}

func (s Service) GetPortfolio(usuario int64) ([]portfolio_domain.Portfolio, error) {
	portfolio, err := s.repository.GetPortfolio(usuario, time.Time{})
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

func (s Service) GetPortfolioByTrimestre(usuario int64, idTrimestre int64) ([]portfolio_domain.Portfolio, error) {
	trimestre, err := s.quarterService.GetQuarter(idTrimestre)
	if err != nil {
		log.Print("Erro ao recuperar trimestre na busca de portfolio.")
		return nil, err
	}

	portfolio, err := s.repository.GetPortfolio(usuario, trimestre.DataFim)
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