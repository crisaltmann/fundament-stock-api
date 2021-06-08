package holding_domain

import (
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
)

type Holdings struct {
	Holdings []Holding
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