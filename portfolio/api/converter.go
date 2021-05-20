package portfolio_api

import (
	portfolio_domain "github.com/crisaltmann/fundament-stock-api/portfolio/domain"
)

func convertDomainsToDtos(itens []portfolio_domain.Portfolio) ([]PortfolioGetResponse, error) {
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

func convertDomainToDto(portfolio portfolio_domain.Portfolio) (PortfolioGetResponse, error) {
	return PortfolioGetResponse{Ativo: Ativo{
		Id:     portfolio.Ativo.Id,
		Codigo: portfolio.Ativo.Codigo,
		Logo:   portfolio.Ativo.Logo,
	}, Quantidade: portfolio.Quantidade, Usuario: portfolio.Usuario}, nil
}