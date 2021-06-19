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