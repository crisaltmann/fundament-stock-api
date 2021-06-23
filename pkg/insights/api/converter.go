package insight_api

import (
	asset_domain "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	insight_domain "github.com/crisaltmann/fundament-stock-api/pkg/insights/domain"
	quarter_domain "github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
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
		Trimestre:    convertQuarterToDto(insight.Trimestre),
		Ativo:        convertAssetToDto(insight.Ativo),
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
		Trimestre:         convertQuarterToDto(summary.Trimestre),
		AtivoMaiorReceita: convertAssetToDto(summary.AtivoMaiorReceita),
		ReceitaMaiorDelta: summary.ReceitaMaiorDelta,
		AtivoMaiorEbitda:  convertAssetToDto(summary.AtivoMaiorEbitda),
		EbitdaMaiorDelta:  summary.EbitdaMaiorDelta,
		AtivoMaiorLucro:   convertAssetToDto(summary.AtivoMaiorLucro),
		LucroMaiorDelta:   summary.LucroMaiorDelta,
		AtivoMaiorDivida:  convertAssetToDto(summary.AtivoMaiorDivida),
		DividaDelta:       summary.DividaDelta,
	}
}

func convertAssetToDto(asset asset_domain.Asset) Ativo {
	return Ativo{
		Id:     asset.Id,
		Codigo: asset.Codigo,
		Nome:   asset.Nome,
	}
}

func convertQuarterToDto(quarter quarter_domain.Trimestre) Trimestre {
	return Trimestre{
		Id:        quarter.Id,
		Ano:       quarter.Ano,
		Trimestre: quarter.Trimestre,
		Codigo:    quarter.Codigo,
	}
}

