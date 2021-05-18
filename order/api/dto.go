package order_api


type OrderPostRequest struct {
	Ativo      int64 	`json:"id_ativo"`
	Quantidade int 		`json:"quantidade"`
	Valor      float32	`json:"valor"`
	Tipo       string	`json:"tipo"`
}
