package holding_service

import (
	holding_domain "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"
)

type Service struct {
//	Repository Repository
}

type Repository interface {
//	GetPortfolio(usuario string) ([]portfolio_domain.Portfolio, error)
}

//func NewService(repository Repository) Service {
//	return Service{Repository: repository}
//}

func NewService() Service {
	return Service{}
}


func (s Service) GetHolding(usuario string) (holding_domain.Holdings, error) {
	holdings := make([]holding_domain.Holding, 0)

	h1 := holding_domain.Holding{
		Id:             1,
		Trimestre:      2,
		Ativo:          3,
		ReceitaLiquida: 100,
		Ebitda:         200,
		MargemEbitda:   0,
		LucroLiquido:   0,
		MargemLiquida:  0,
		DividaLiquida:  0,
		DivEbitda:      0,
	}

	h2 := holding_domain.Holding{
		Id:             2,
		Trimestre:      2,
		Ativo:          2,
		ReceitaLiquida: 100,
		Ebitda:         200,
		MargemEbitda:   2222,
		LucroLiquido:   0,
		MargemLiquida:  0,
		DividaLiquida:  0123111111,
		DivEbitda:      0,
	}

	holdings = append(holdings, h1)
	holdings = append(holdings, h2)

	return holding_domain.Holdings{Holdings: holdings}, nil
}