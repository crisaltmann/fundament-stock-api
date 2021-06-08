package holding_api

import "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"

func convertHoldingsDomainToDto(holdings holding_domain.Holdings) Holdings {
	holdingsDTO := make([]Holding, 0)
	if len(holdings.Holdings) > 0 {
		for _, h := range holdings.Holdings {
			holdingsDTO = append(holdingsDTO, convertDomainToDto(h))
		}
	}
	return Holdings{Holdings: holdingsDTO}
}

func convertDomainToDto(holding holding_domain.Holding) Holding {
	return Holding{
		Trimestre:      convertTrimestreToDto(holding),
		ReceitaLiquida: holding.ReceitaLiquida,
		Ebitda:         holding.Ebitda,
		MargemEbitda:   holding.MargemEbitda,
		LucroLiquido:   holding.LucroLiquido,
		MargemLiquida:  holding.MargemLiquida,
		DividaLiquida:  holding.DividaLiquida,
		DivEbitda:      holding.DivEbitda,
	}
}

func convertTrimestreToDto(holding holding_domain.Holding) Trimestre {
	return Trimestre{
		Id:        holding.Trimestre.Id,
		Ano:       holding.Trimestre.Ano,
		Trimestre: holding.Trimestre.Trimestre,
	}
}
