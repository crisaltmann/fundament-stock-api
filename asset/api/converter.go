package api

import (
	"github.com/crisaltmann/fundament-stock-api/asset/service"
)

func convertToDtos(assets []service.Asset) []Asset {
	assetDtos := make([]Asset, 10)
	for _, asset := range assets {
		assetDtos = append(assetDtos, convertToDto(asset))
	}
	return assetDtos
}

func convertToDto(asset service.Asset) Asset {
	return Asset{asset.Id, asset.Codigo, asset.Nome}
}
