package holding_api

type Holdings struct {
	Holdings 		[]Holding 		`json:"holdings"`
	Consolidated	[]AnnualHolding `json:"consolidated"`
}

type Holding struct {
	Trimestre    		Trimestre `json:"trimestre"`
	ReceitaLiquida	    int64 `json:"receita_liquida"`
	Ebitda	     		int64 `json:"ebitda"`
	MargemEbitda		float32 `json:"margem_ebitda"`
	LucroLiquido	    int64 `json:"lucro_liquido"`
	MargemLiquida		float32 `json:"margem_liquida"`
	DividaLiquida	    int64 `json:"divida_liquida"`
	DivEbitda			float32 `json:"div_ebitda"`
	HoldingsAtivo		[]HoldingAtivo `json:"ativos,omitempty"`
}

type HoldingAtivo struct {
	Ativo				Ativo 		`json:"ativo"`
	Trimestre    		int64		`json:"id_trimestre"`
	ReceitaLiquida	    int64 		`json:"receita_liquida"`
	Ebitda	     		int64		`json:"ebitda"`
	MargemEbitda		float32     `json:"margem_ebitda"`
	LucroLiquido	    int64       `json:"lucro_liquido"`
	MargemLiquida		float32     `json:"margem_liquida"`
	DividaLiquida	    int64       `json:"divida_liquida"`
	DivEbitda			float32     `json:"div_ebitda"`
}

type Ativo struct {
	Id     	int64  `json:"id"`
	Codigo 	string `json:"codigo"`
	Nome   	string `json:"nome"`
}

type Trimestre struct {
	Id			int64			`json:"id"`
	Ano			int				`json:"ano"`
	Trimestre   int				`json:"trimestre"`
}

type AnnualHoldings struct {
	Consolidated		[]AnnualHolding		`json:"consolidated"`
}

type AnnualHolding struct {
	Ano					int64				`json:"ano"`
	ReceitaLiquida	    int64				`json:"receita_liquida"`
	Ebitda	     		int64				`json:"ebitda"`
	MargemEbitda		float32				`json:"margem_ebitda"`
	LucroLiquido	    int64				`json:"lucro_liquido"`
	MargemLiquida		float32				`json:"margem_liquida"`
	DividaLiquida	    int64				`json:"divida_liquida"`
	DivEbitda			float32				`json:"div_ebitda"`
}