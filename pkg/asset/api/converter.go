package asset_api

import (
	"fmt"
	asset_domain2 "github.com/crisaltmann/fundament-stock-api/pkg/asset/domain"
	"strconv"
)

func convertToDtos(assets []asset_domain2.Asset) []AssetResponse {
	assetDtos := make([]AssetResponse, 0)
	for _, asset := range assets {
		assetDtos = append(assetDtos, convertDomainToDto(asset))
	}
	return assetDtos
}

func convertDomainToDto(asset asset_domain2.Asset) AssetResponse {
	return AssetResponse{asset.Id, asset.Codigo, asset.Nome, asset.Logo}
}

func convertPostRequestToDomain(request AssetPostRequest) asset_domain2.Asset {
	return asset_domain2.Asset{Codigo: request.Codigo, Nome: request.Nome, Logo: request.Logo}
}

func convertPutRequestToDomain(request AssetPutRequest, id string) (asset_domain2.Asset, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return asset_domain2.Asset{}, fmt.Errorf("Id informado é inválido.")
	}
	return asset_domain2.Asset{Id: idInt, Codigo: request.Codigo, Nome: request.Nome, Logo: request.Logo}, nil
}
