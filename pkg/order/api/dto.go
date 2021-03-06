package order_api

import (
	"time"
)

type OrderPostRequest struct {
	Ativo      int64 		`json:"id_ativo"`
	Quantidade int 			`json:"quantidade"`
	Valor      float32		`json:"valor"`
	Tipo       string		`json:"tipo"`
	Data       time.Time 	`json:"data"`
	IdUsuario  int64		`json:"id_usuario"`
}

type OrderGetResponse struct {
	Id         int64        `json:"id"`
	Ativo      int64 		`json:"id_ativo"`
	Quantidade int 			`json:"quantidade"`
	Valor      float32		`json:"valor"`
	Tipo       string		`json:"tipo"`
	Data       time.Time 	`json:"data"`
	IdUsuario  int64		`json:"id_usuario"`
}
