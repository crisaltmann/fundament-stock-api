package asset_api

import (
	"fmt"
	asset_domain "github.com/crisaltmann/fundament-stock-api/asset/domain"
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
	return AssetResponse{asset.Id, asset.Codigo, asset.Nome, asset.Logo}
}

func convertPostRequestToDomain(request AssetPostRequest) asset_domain.Asset {
	return asset_domain.Asset{Codigo: request.Codigo, Nome: request.Nome, Logo: request.Logo}
}

func convertPutRequestToDomain(request AssetPutRequest, id string) (asset_domain.Asset, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return asset_domain.Asset{}, fmt.Errorf("Id informado é inválido.")
	}
	return asset_domain.Asset{Id: idInt, Codigo: request.Codigo, Nome: request.Nome, Logo: request.Logo}, nil
}
