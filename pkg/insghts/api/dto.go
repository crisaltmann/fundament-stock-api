package insight_api

//type TrimestreGetResponse struct {
//	Id			int64			`json:"id"`
//	Codigo		string			`json:"codigo"`
//	Ano			int				`json:"ano"`
//	Trimestre   int				`json:"trimestre"`
//	DataInicio	time.Time		`json:"data_inicio"`
//	DataFim		time.Time		`json:"data_fim"`
//}

type Insight struct {
	Id			    int64		`json:"id"`
	IdTrimestre     int64		`json:"trimestre"`
	ReceitaAtivo    int64		`json:"receita_ativo"`
	ReceitaDelta	float32		`json:"receita_delta"`
}