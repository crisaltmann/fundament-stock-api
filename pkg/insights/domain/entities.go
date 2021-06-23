package insight_domain

import (
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
)

type Insight struct {
	Id			    int64
	Usuario		    int64
	Trimestre     	quarter_domain.Trimestre
	IdTrimestre		int64
	Ativo			asset_domain.Asset
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
	IdTrimestre				int64
	Trimestre     			quarter_domain.Trimestre

	AtivoMaiorReceita    	asset_domain.Asset
	ReceitaMaiorDelta		float32

	AtivoMaiorEbitda		asset_domain.Asset
	EbitdaMaiorDelta		float32

	AtivoMaiorLucro			asset_domain.Asset
	LucroMaiorDelta			float32

	AtivoMaiorDivida		asset_domain.Asset
	DividaDelta				float32
}

func (s *InsightSummary) ToStruct() InsightSummary {
	return InsightSummary{
		IdTrimestre: 	   s.IdTrimestre,
		Trimestre:         s.Trimestre,
		AtivoMaiorReceita: s.AtivoMaiorReceita,
		ReceitaMaiorDelta: s.ReceitaMaiorDelta,
		AtivoMaiorEbitda:  s.AtivoMaiorEbitda,
		EbitdaMaiorDelta:  s.EbitdaMaiorDelta,
		AtivoMaiorLucro:   s.AtivoMaiorLucro,
		LucroMaiorDelta:   s.LucroMaiorDelta,
		AtivoMaiorDivida:  s.AtivoMaiorDivida,
		DividaDelta:       s.DividaDelta,
	}

}