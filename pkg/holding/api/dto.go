package holding_api

type Holdings struct {
	Holdings []Holding `json:"holdings"`
}

type Holding struct {
	Id              	int64 `json:"id"`
	Trimestre    		int64 `json:"id_trimestre"`
	ReceitaLiquida	    int64 `json:"receita_liquida"`
	Ebitda	     		int64 `json:"ebitda"`
	MargemEbitda		float32 `json:"margem_ebitda"`
	LucroLiquido	    int64 `json:"lucro_liquido"`
	MargemLiquida		float32 `json:"margem_liquida"`
	DividaLiquida	    int64 `json:"divida_liquida"`
	DivEbitda			float32 `json:"div_ebitda"`
}
