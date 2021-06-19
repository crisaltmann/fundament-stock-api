package quarter_domain

import "time"

type Trimestre struct {
	Id					int64
	Codigo				string
	Ano					int
	Trimestre   		int
	DataInicio			time.Time
	DataFim				time.Time
	TrimestreAnterior	int64
}