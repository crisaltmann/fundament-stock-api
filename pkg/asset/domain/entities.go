package asset_domain

import "time"

type Asset struct {
	Id     	int64
	Codigo 	string
	Nome   	string
	Logo   	string
	Cotacao float32
}

type AssetPrice struct {
	Id 		int64
	Ativo   int64
	Cotacao float32
	Data    time.Time
}

type AssetQuarterlyResult struct {
	Id              	int64
	Trimestre    		int64
	Ativo	     		int64
	ReceitaLiquida	    int64
	Ebitda	     		int64
	MargemEbitda		float32
	LucroLiquido	    int64
	MargemLiquida		float32
	DividaLiquida	    int64
	DivEbitda			float32
}
