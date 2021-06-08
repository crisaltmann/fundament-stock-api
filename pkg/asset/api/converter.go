package asset_api

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"strconv"
)

func convertToDtos(assets []asset_domain.Asset) []AssetResponse {
	assetDtos := make([]AssetResponse, 0)
	for _, asset := range assets {
		assetDtos = append(assetDtos, convertDomainToDto(asset))
	}
	return assetDtos
}

func convertDomainToDto(asset asset_domain.Asset) AssetResponse {
	return AssetResponse{asset.Id, asset.Codigo, asset.Nome, asset.Logo, asset.Cotacao, asset.Total}
}

func convertPostRequestToDomain(request AssetPostRequest) asset_domain.Asset {
	return asset_domain.Asset{Codigo: request.Codigo, Nome: request.Nome, Logo: request.Logo, Total: request.Total}
}

func convertPutRequestToDomain(request AssetPutRequest, id string) (asset_domain.Asset, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return asset_domain.Asset{}, fmt.Errorf("Id informado é inválido.")
	}
	return asset_domain.Asset{Id: idInt, Codigo: request.Codigo, Nome: request.Nome, Logo: request.Logo,
		Cotacao: request.Cotacao, Total: request.Total}, nil
}

func convertPostQuarterlyRequestToDomain(request QuarterlyResultPostRequest) asset_domain.AssetQuarterlyResult {
	return asset_domain.AssetQuarterlyResult{
		Trimestre:      request.Trimestre,
		Ativo:          request.Ativo,
		ReceitaLiquida: request.ReceitaLiquida,
		Ebitda:         request.Ebitda,
		LucroLiquido:   request.LucroLiquido,
		DividaLiquida:  request.DividaLiquida,
	}
}

func convertQuarterlyResultToDtos(quarterResult asset_domain.AssetQuarterlyResult) QuarterlyResultResponse {
	return QuarterlyResultResponse{
		Id:             quarterResult.Id,
		Trimestre:      quarterResult.Trimestre,
		Ativo:          quarterResult.Ativo,
		ReceitaLiquida: quarterResult.ReceitaLiquida,
		Ebitda:         quarterResult.Ebitda,
		LucroLiquido:   quarterResult.LucroLiquido,
		DividaLiquida:  quarterResult.DividaLiquida,
		MargemEbitda:   quarterResult.MargemEbitda,
		MargemLiquida:  quarterResult.MargemLiquida,
		DivEbitda:      quarterResult.DivEbitda,
	}
}

func convertQuarterlyResultsToDtos(quarterResults []asset_domain.AssetQuarterlyResult) []QuarterlyResultResponse {
	quarterlyDtos := make([]QuarterlyResultResponse, 0)
	for _, qra := range quarterResults {
		quarterlyDtos = append(quarterlyDtos, convertQuarterlyResultToDtos(qra))
	}
	return quarterlyDtos
}