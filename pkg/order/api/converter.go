package order_api

import (
	"fmt"
	order_domain2 "github.com/crisaltmann/fundament-stock-api/pkg/order/domain"
)

func convertPostRequestToDomain(request OrderPostRequest) (order_domain2.Order, error) {
	orderType, err := order_domain2.ConvertType(request.Tipo)
	if err != nil {
		return order_domain2.Order{}, fmt.Errorf("Erro ao converter tipo de order.")
	}
	return order_domain2.Order{Ativo: request.Ativo, Quantidade: request.Quantidade,
		Valor: request.Valor, Tipo: orderType, Data: request.Data, Usuario: request.IdUsuario}, nil
}

func convertDomainsToDtos(orders []order_domain2.Order) ([]OrderGetResponse, error) {
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

func convertDomainToDto(order order_domain2.Order) (OrderGetResponse, error) {
	orderType, err := order_domain2.ConvertTypeToString(order.Tipo)
	if err != nil {
		return OrderGetResponse{}, fmt.Errorf("Erro ao converter tipo de order.")
	}
	return OrderGetResponse{Id: order.Id, Ativo: order.Ativo, Quantidade: order.Quantidade, Valor: order.Valor,
		Tipo: orderType, Data: order.Data, IdUsuario: order.Usuario}, nil
}