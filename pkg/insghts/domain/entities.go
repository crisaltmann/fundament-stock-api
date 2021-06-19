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

type InsightsSummary struct {
	Insights 		[]InsightSummary
}

type InsightSummary struct {
	Trimestre 				int64

	AtivoMaiorReceita    	int64
	ReceitaMaiorDelta		float32

	AtivoMaiorEbitda		int64
	EbitdaMaiorDelta		float32

	AtivoMaiorLucro			int64
	LucroMaiorDelta			float32

	AtivoMaiorDivida		int64
	DividaDelta				float32
}