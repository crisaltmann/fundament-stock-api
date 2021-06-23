package insight_api

import (
	insight_domain "github.com/crisaltmann/fundament-stock-api/pkg/insights/domain"
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

func convertInsightsSummaryDomainToDto(insights insight_domain.InsightsSummary) InsightsSummary {
	insightsDTO := make([]InsightSummary, 0)
	if len(insights.Insights) > 0 {
		for _, h := range insights.Insights {
			insightsDTO = append(insightsDTO, convertSummaryDomainToDto(h))
		}
	}
	return InsightsSummary{
		Insights: insightsDTO,
	}
}

func convertSummaryDomainToDto(summary insight_domain.InsightSummary) InsightSummary {
	return InsightSummary{
		Trimestre:         summary.Trimestre,
		AtivoMaiorReceita: summary.AtivoMaiorReceita,
		ReceitaMaiorDelta: summary.ReceitaMaiorDelta,
		AtivoMaiorEbitda:  summary.AtivoMaiorEbitda,
		EbitdaMaiorDelta:  summary.EbitdaMaiorDelta,
		AtivoMaiorLucro:   summary.AtivoMaiorLucro,
		LucroMaiorDelta:   summary.LucroMaiorDelta,
		AtivoMaiorDivida:  summary.AtivoMaiorDivida,
		DividaDelta:       summary.DividaDelta,
	}
}

