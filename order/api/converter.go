package order_api

import order_domain "github.com/crisaltmann/fundament-stock-api/order/domain"

//func convertToDtos(assets []domain.Asset) []AssetResponse {
//	assetDtos := make([]AssetResponse, 0)
//	for _, asset := range assets {
//		assetDtos = append(assetDtos, convertDomainToDto(asset))
//	}
//	return assetDtos
//}
//
//func convertDomainToDto(asset domain.Asset) AssetResponse {
//	return AssetResponse{asset.Id, asset.Codigo, asset.Nome, asset.Logo}
//}
//
func convertPostRequestToDomain(request OrderPostRequest) order_domain.Order {
	return order_domain.Order{Ativo: request.Ativo, Quantidade: request.Quantidade, Valor: request.Valor}
}
//
//func convertPutRequestToDomain(request AssetPutRequest, id string) (domain.Asset, error) {
//	idInt, err := strconv.ParseInt(id, 10, 64)
//	if err != nil {
//		return domain.Asset{}, fmt.Errorf("Id informado é inválido.")
//	}
//	return domain.Asset{Id: idInt, Codigo: request.Codigo, Nome: request.Nome, Logo: request.Logo}, nil
//}
