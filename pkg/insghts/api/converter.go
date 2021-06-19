package insight_api

import (
	insight_domain "github.com/crisaltmann/fundament-stock-api/pkg/insghts/domain"
)

func convertInsightsDomainToDto(insights []insight_domain.Insight) []Insight {
	insightsDTO := make([]Insight, 0)
	if len(insights) > 0 {
		for _, h := range insights {
			insightsDTO = append(insightsDTO, convertDomainToDto(h))
		}
	}
	return insightsDTO
}

func convertDomainToDto(insight insight_domain.Insight) Insight {
	return Insight{
		Id:           insight.Id,
		Usuario:      insight.Usuario,
		IdTrimestre:  insight.IdTrimestre,
		IdAtivo:      insight.IdAtivo,
		ReceitaDelta: insight.ReceitaDelta,
		EbitdaDelta:  insight.EbitdaDelta,
		LucroDelta:   insight.LucroDelta,
		DividaDelta:  insight.DividaDelta,
	}
}


