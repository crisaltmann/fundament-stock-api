package insight_domain

type Insight struct {
	Id			    int64
	Usuario		    int64
	IdTrimestre     int64
	IdAtivo			int64
	ReceitaDelta	float32
	EbitdaDelta		float32
	LucroDelta		float32
	DividaDelta		float32
}