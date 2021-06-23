package insight_api

type Insight struct {
	Id			    int64		`json:"id"`
	Usuario			int64		`json:"id_usuario"`
	IdTrimestre     int64		`json:"trimestre"`
	IdAtivo			int64		`json:"id_ativo"`
	ReceitaDelta	float32		`json:"receita_delta"`
	EbitdaDelta		float32		`json:"ebitda_delta"`
	LucroDelta		float32		`json:"lucro_delta"`
	DividaDelta		float32		`json:"divida_delta"`
}

type InsightsSummary struct {
	Insights 		[]InsightSummary	`json:"insights"`
}

type InsightSummary struct {
	Trimestre 				int64		`json:"trimestre"`

	AtivoMaiorReceita    	int64		`json:"ativo_maior_receita"`
	ReceitaMaiorDelta		float32		`json:"receita_maior_delta"`

	AtivoMaiorEbitda		int64		`json:"ativo_maior_ebitda"`
	EbitdaMaiorDelta		float32		`json:"ebitda_maior_delta"`

	AtivoMaiorLucro			int64		`json:"ativo_maior_lucro"`
	LucroMaiorDelta			float32		`json:"lucro_maior_delta"`

	AtivoMaiorDivida		int64		`json:"ativo_maior_divida"`
	DividaDelta				float32		`json:"divida_delta"`
}