package quarter_api

import (
	"github.com/crisaltmann/fundament-stock-api/pkg/quarter/domain"
)

func convertToDtos(quarters []quarter_domain.Trimestre) []TrimestreGetResponse {
	quartersDtos := make([]TrimestreGetResponse, 0)
	for _, quarter := range quarters {
		quartersDtos = append(quartersDtos, convertDomainToDto(quarter))
	}
	return quartersDtos
}

func convertDomainToDto(quarter quarter_domain.Trimestre) TrimestreGetResponse {
	return TrimestreGetResponse{
		Id:         quarter.Id,
		Codigo:     quarter.Codigo,
		Ano:        quarter.Ano,
		Trimestre:  quarter.Trimestre,
		DataInicio: quarter.DataInicio,
		DataFim:    quarter.DataFim,
		TrimestreAnterior: quarter.TrimestreAnterior,
	}
}