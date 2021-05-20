package portfolio_api

import (
	portfolio_domain2 "github.com/crisaltmann/fundament-stock-api/pkg/portfolio/domain"
)

func convertDomainsToDtos(itens []portfolio_domain2.Portfolio) ([]PortfolioGetResponse, error) {
	portfolioDtos := make([]PortfolioGetResponse, 0)
	for _, item := range itens {
		dto, err := convertDomainToDto(item)
		if err != nil {
			return portfolioDtos, err
		}
		portfolioDtos = append(portfolioDtos, dto)
	}
	return portfolioDtos, nil
}

func convertDomainToDto(portfolio portfolio_domain2.Portfolio) (PortfolioGetResponse, error) {
	return PortfolioGetResponse{Ativo: Ativo{
		Id:     portfolio.Ativo.Id,
		Codigo: portfolio.Ativo.Codigo,
		Logo:   portfolio.Ativo.Logo,
	}, Quantidade: portfolio.Quantidade, Usuario: portfolio.Usuario}, nil
}