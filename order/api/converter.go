package order_api

import (
	"fmt"
	order_domain "github.com/crisaltmann/fundament-stock-api/order/domain"
)

func convertPostRequestToDomain(request OrderPostRequest) (order_domain.Order, error) {
	orderType, err := order_domain.ConvertType(request.Tipo)
	if err != nil {
		return order_domain.Order{}, fmt.Errorf("Erro ao converter tipo de order.")
	}
	return order_domain.Order{Ativo: request.Ativo, Quantidade: request.Quantidade, Valor: request.Valor, Tipo: orderType}, nil
}