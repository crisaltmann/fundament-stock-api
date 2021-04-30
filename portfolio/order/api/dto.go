package orderapi

type OrderPostRequest struct {
	IdAtivo int64 `json:"id_ativo"`
	Data int64 `json:"data"`
	Quantidade int64 `json:"quantidade"`
	Valor int64 `json:"valor"`
}