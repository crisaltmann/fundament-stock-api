package api

import (
	"github.com/crisaltmann/fundament-stock-api/asset/domain"
)

func convertToDtos(assets []domain.Asset) []Asset {
	assetDtos := make([]Asset, 10)
	for _, asset := range assets {
		assetDtos = append(assetDtos, convertToDto(asset))
	}
	return assetDtos
}

func convertToDto(asset domain.Asset) Asset {
	return Asset{asset.Id, asset.Codigo, asset.Nome}
}
