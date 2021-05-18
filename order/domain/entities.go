package order_domain

import "fmt"

type OrderType string

const (
	BuyOrder  OrderType = "Compra"
	SellOrder OrderType = "Venda"
)

type Order struct {
	Id         int64
	Ativo      int64
	Quantidade int
	Valor      float32
	Tipo       OrderType
}

func ConvertType(orderType string) (OrderType, error) {
	switch orderType {
	case "C":
		return BuyOrder, nil
	case "V":
		return SellOrder, nil
	default:
		return "", fmt.Errorf("")
	}
}
