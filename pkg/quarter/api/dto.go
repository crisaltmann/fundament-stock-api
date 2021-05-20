package quarter_api

import "time"

type TrimestreGetResponse struct {
	Id			int64			`json:"id"`
	Codigo		string			`json:"codigo"`
	Ano			int				`json:"ano"`
	Trimestre   int				`json:"trimestre"`
	DataInicio	time.Time		`json:"data_inicio"`
	DataFim		time.Time		`json:"data_fim"`
}
