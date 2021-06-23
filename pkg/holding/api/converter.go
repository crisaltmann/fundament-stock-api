package holding_api

import "github.com/crisaltmann/fundament-stock-api/pkg/holding/domain"

func convertHoldingsDomainToDto(holdings holding_domain.Holdings, expandirAtivos bool) Holdings {
	holdingsDTO := make([]Holding, 0)
	if len(holdings.Holdings) > 0 {
		for _, h := range holdings.Holdings {
			holdingsDTO = append(holdingsDTO, convertDomainToDto(h, expandirAtivos))
		}
	}

	consolidated := convertAnnualHoldingsToDto(holdings.Consolidated)

	return Holdings{Holdings: holdingsDTO, Consolidated: consolidated}
}

func convertDomainToDto(holding holding_domain.Holding, expandirAtivos bool) Holding {
	holdingReturn := Holding{
		Trimestre:      convertTrimestreToDto(holding),
		ReceitaLiquida: holding.ReceitaLiquida,
		Ebitda:         holding.Ebitda,
		MargemEbitda:   holding.MargemEbitda,
		LucroLiquido:   holding.LucroLiquido,
		MargemLiquida:  holding.MargemLiquida,
		DividaLiquida:  holding.DividaLiquida,
		DivEbitda:      holding.DivEbitda,
	}

	if expandirAtivos {
		holdingReturn.HoldingsAtivo = convertHoldingAtivosToDto(holding.HoldingsAtivo)
	}

	return holdingReturn
}

func convertHoldingAtivosToDto(ativos []holding_domain.HoldingAtivo) []HoldingAtivo{
	 if len(ativos) == 0 {
	 	return nil
	 }
	 ativosReturn := make([]HoldingAtivo, 0)
	 for _, ativo := range ativos {
	 	ativosReturn = append(ativosReturn, convertHoldingAtivoToDto(ativo))
	 }
	 return ativosReturn
}

func convertHoldingAtivoToDto(ativo holding_domain.HoldingAtivo) HoldingAtivo {
	return HoldingAtivo{
		Ativo:  		convertAtivoToDto(ativo),
		Trimestre:      ativo.Trimestre,
		ReceitaLiquida: ativo.ReceitaLiquida,
		Ebitda:         ativo.Ebitda,
		MargemEbitda:   ativo.MargemEbitda,
		LucroLiquido:   ativo.LucroLiquido,
		MargemLiquida:  ativo.MargemLiquida,
		DividaLiquida:  ativo.DividaLiquida,
		DivEbitda:      ativo.DivEbitda,
	}
}

func convertTrimestreToDto(holding holding_domain.Holding) Trimestre {
	return Trimestre{
		Id:        holding.Trimestre.Id,
		Ano:       holding.Trimestre.Ano,
		Trimestre: holding.Trimestre.Trimestre,
	}
}

func convertAtivoToDto(holding holding_domain.HoldingAtivo) Ativo {
	return Ativo{
		Id:     holding.Ativo.Id,
		Codigo: holding.Ativo.Codigo,
		Nome:   holding.Ativo.Nome,
	}
}

func convertAnnualHoldingsToDto(consolidated holding_domain.AnnualHoldings) AnnualHoldings {
	dto := make([]AnnualHolding, 0)
	for _, annual := range consolidated.Consolidated {
		dto = append(dto, convertAnnualHoldingToDto(annual))
	}
	return AnnualHoldings{Consolidated: dto}
}

func convertAnnualHoldingToDto(annual holding_domain.AnnualHolding) AnnualHolding {
	return AnnualHolding{
		Ano:            annual.Ano,
		ReceitaLiquida: annual.ReceitaLiquida,
		Ebitda:         annual.Ebitda,
		MargemEbitda:   annual.MargemEbitda,
		LucroLiquido:   annual.LucroLiquido,
		MargemLiquida:  annual.MargemLiquida,
		DividaLiquida:  annual.DividaLiquida,
		DivEbitda:      annual.DivEbitda,
	}
}