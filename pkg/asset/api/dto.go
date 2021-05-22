package asset_api

type AssetResponse struct {
	Id     	int64  `json:"id"`
	Codigo 	string `json:"codigo"`
	Nome   	string `json:"nome"`
	Logo   	string `json:"logo"`
	Cotacao float32 `json:"cotacao"`
}

type AssetPostRequest struct {
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`
	Logo   string `json:"logo""`
}

type AssetPutRequest struct {
	Codigo  string `json:"codigo"`
	Nome    string `json:"nome"`
	Logo    string `json:"logo"`
	Cotacao float32 `json:"cotacao"`
}

type QuarterlyResultPostRequest struct {
	Trimestre    		int64	`json:"id_trimestre"`
	Ativo	     		int64	`json:"id_ativo"`
	ReceitaLiquida	    int64	`json:"receita_liquida"`
	Ebitda	     		int64	`json:"ebitda"`
	LucroLiquido	    int64	`json:"lucro_liquido"`
	DividaLiquida	    int64	`json:"divida_liquida"`
}

type QuarterlyResultResponse struct {
	Id              	int64   `json:"id"`
	Trimestre    		int64	`json:"id_trimestre"`
	Ativo	     		int64	`json:"id_ativo"`
	ReceitaLiquida	    int64	`json:"receita_liquida"`
	Ebitda	     		int64	`json:"ebitda"`
	LucroLiquido	    int64	`json:"lucro_liquido"`
	DividaLiquida	    int64	`json:"divida_liquida"`
	MargemEbitda		float32 `json:"margem_ebitda"`
	MargemLiquida		float32 `json:"margem_liquida"`
	DivEbitda			float32 `json:"divida_ebitda"`
}