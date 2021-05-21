package order_api

import (
	"fmt"
	"github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
)

func convertPostRequestToDomain(request OrderPostRequest) (order_domain.Order, error) {
	orderType, err := order_domain.ConvertType(request.Tipo)
	if err != nil {
		return order_domain.Order{}, fmt.Errorf("Erro ao converter tipo de order.")
	}
	return order_domain.Order{Ativo: request.Ativo, Quantidade: request.Quantidade,
		Valor: request.Valor, Tipo: orderType, Data: request.Data, Usuario: request.IdUsuario}, nil
}

func convertDomainsToDtos(orders []order_domain.Order) ([]OrderGetResponse, error) {
	orderDtos := make([]OrderGetResponse, 0)
	for _, order := range orders {
		dto, err := convertDomainToDto(order)
		if err != nil {
			return orderDtos, err
		}
		orderDtos = append(orderDtos, dto)
	}
	return orderDtos, nil
}

func convertDomainToDto(order order_domain.Order) (OrderGetResponse, error) {
	orderType, err := order_domain.ConvertTypeToString(order.Tipo)
	if err != nil {
		return OrderGetResponse{}, fmt.Errorf("Erro ao converter tipo de order.")
	}
	return OrderGetResponse{Id: order.Id, Ativo: order.Ativo, Quantidade: order.Quantidade, Valor: order.Valor,
		Tipo: orderType, Data: order.Data, IdUsuario: order.Usuario}, nil
}