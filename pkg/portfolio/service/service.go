package portfolio_service

import (
	"fmt"
	portfolio_domain2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
	portfolio_repository2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/repository"
	"strconv"
)

type Service struct {
	Repository *portfolio_repository2.Repository
}

func (s Service) GetPortfolio(usuario string) ([]portfolio_domain2.Portfolio, error) {
	portfolio, err := s.Repository.GetPortfolio(usuario)
	if err != nil {
		return portfolio, err
	}
	for idx, item := range portfolio {
		if item.Ativo.Cotacao > 0 {
			portfolio[idx].Valor = roundFloat(float32(item.Quantidade) * item.Ativo.Cotacao, 2)
		}
	}
	return portfolio, nil
}

func roundFloat(x float32, prec int) float32 {
	i := fmt.Sprintf("%.2f", x)
	f, _ := strconv.ParseFloat(i, 2)
	return float32(f)
}