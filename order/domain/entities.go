package order_domain

import (
	"fmt"
	"time"
)

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
	Data	   time.Time
	Usuario	   int64
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

func ConvertTypeToString(orderType OrderType) (string, error) {
	switch orderType {
	case BuyOrder:
		return "B", nil
	case SellOrder:
		return "S", nil
	default:
		return "", fmt.Errorf("")
	}
}
