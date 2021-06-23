package insight_api

type Insight struct {
	Id			    int64		`json:"id"`
	Usuario			int64		`json:"id_usuario"`
	Trimestre		Trimestre	`json:"trimestre"`
	Ativo			Ativo		`json:"ativo"`
	ReceitaDelta	float32		`json:"receita_delta"`
	EbitdaDelta		float32		`json:"ebitda_delta"`
	LucroDelta		float32		`json:"lucro_delta"`
	DividaDelta		float32		`json:"divida_delta"`
}

type InsightsSummary struct {
	Insights 		[]InsightSummary	`json:"insights"`
}

type InsightSummary struct {
	Trimestre 				Trimestre		`json:"trimestre"`

	AtivoMaiorReceita    	Ativo		`json:"ativo_maior_receita"`
	ReceitaMaiorDelta		float32		`json:"receita_maior_delta"`

	AtivoMaiorEbitda		Ativo		`json:"ativo_maior_ebitda"`
	EbitdaMaiorDelta		float32		`json:"ebitda_maior_delta"`

	AtivoMaiorLucro			Ativo		`json:"ativo_maior_lucro"`
	LucroMaiorDelta			float32		`json:"lucro_maior_delta"`

	AtivoMaiorDivida		Ativo		`json:"ativo_maior_divida"`
	DividaDelta				float32		`json:"divida_delta"`
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
	Codigo		string			`json:"codigo"`
}
