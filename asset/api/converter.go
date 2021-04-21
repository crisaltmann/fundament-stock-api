package api

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/asset/domain"
	"strconv"
)

func convertToDtos(assets []domain.Asset) []AssetResponse {
	assetDtos := make([]AssetResponse, 0)
	for _, asset := range assets {
		assetDtos = append(assetDtos, convertToDto(asset))
	}
	return assetDtos
}

func convertToDto(asset domain.Asset) AssetResponse {
	return AssetResponse{asset.Id, asset.Codigo, asset.Nome}
}

func convertPostRequestToDomain(request AssetPostRequest) domain.Asset {
	return domain.Asset{Codigo: request.Codigo, Nome: request.Nome}
}

func convertPutRequestToDomain(request AssetPutRequest, id string) (domain.Asset, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return domain.Asset{}, fmt.Errorf("Id informado é inválido.")
	}
	return domain.Asset{Id: idInt, Codigo: request.Codigo, Nome: request.Nome}, nil
}
