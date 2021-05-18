package order_api

import order_domain "github.com/crisaltmann/fundament-stock-api/order/domain"

func convertPostRequestToDomain(request OrderPostRequest) order_domain.Order {
	return order_domain.Order{Ativo: request.Ativo, Quantidade: request.Quantidade, Valor: request.Valor}
}