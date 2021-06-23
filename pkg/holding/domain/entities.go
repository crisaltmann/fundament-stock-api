package holding_domain

import (
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
)

type Holdings struct {
	Holdings []Holding
	Consolidated AnnualHoldings
}

type AnnualHoldings struct {
	Consolidated		[]AnnualHolding
}

type AnnualHolding struct {
	Ano					int64
	ReceitaLiquida	    int64
	Ebitda	     		int64
	MargemEbitda		float32
	LucroLiquido	    int64
	MargemLiquida		float32
	DividaLiquida	    int64
	DivEbitda			float32
}

type Holding struct {
	Trimestre    		quarter_domain.Trimestre
	ReceitaLiquida	    int64
	Ebitda	     		int64
	MargemEbitda		float32
	LucroLiquido	    int64
	MargemLiquida		float32
	DividaLiquida	    int64
	DivEbitda			float32
	HoldingsAtivo	    []HoldingAtivo
}

type HoldingAtivo struct {
	Id					int64
	Usuario				int64
	Ativo				asset_domain.Asset
	Trimestre    		int64
	ReceitaLiquida	    int64
	Ebitda	     		int64
	MargemEbitda		float32
	LucroLiquido	    int64
	MargemLiquida		float32
	DividaLiquida	    int64
	DivEbitda			float32
}

func (h *HoldingAtivo) ToStruct() HoldingAtivo {
	return HoldingAtivo{
		Ativo:          h.Ativo,
		Usuario: 		h.Usuario,
		Trimestre:      h.Trimestre,
		ReceitaLiquida: h.ReceitaLiquida,
		Ebitda:         h.Ebitda,
		MargemEbitda:   h.MargemEbitda,
		LucroLiquido:   h.LucroLiquido,
		MargemLiquida:  h.MargemLiquida,
		DividaLiquida:  h.DividaLiquida,
		DivEbitda:      h.DivEbitda,
	}
}

func (h *Holding) ToStruct() Holding {
	return Holding{
		Trimestre:      h.Trimestre,
		ReceitaLiquida: h.ReceitaLiquida,
		Ebitda:         h.Ebitda,
		MargemEbitda:   h.MargemEbitda,
		LucroLiquido:   h.LucroLiquido,
		MargemLiquida:  h.MargemLiquida,
		DividaLiquida:  h.DividaLiquida,
		DivEbitda:      h.DivEbitda,
		HoldingsAtivo:  h.HoldingsAtivo,
	}
}

func (h *AnnualHolding) ToStruct() AnnualHolding {
	return AnnualHolding{
		Ano:            h.Ano,
		ReceitaLiquida: h.ReceitaLiquida,
		Ebitda:         h.Ebitda,
		MargemEbitda:   h.MargemEbitda,
		LucroLiquido:   h.LucroLiquido,
		MargemLiquida:  h.MargemLiquida,
		DividaLiquida:  h.DividaLiquida,
		DivEbitda:      h.DivEbitda,
	}
}