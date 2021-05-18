package order_api

//type AssetResponse struct {
//	Id     int64  `json:"id"`
//	Codigo string `json:"codigo"`
//	Nome   string `json:"nome"`
//	Logo   string `json:"logo"`
//}
//
//type AssetPostRequest struct {
//	Codigo string `json:"codigo"`
//	Nome   string `json:"nome"`
//	Logo   string `json:"logo""`
//}
//
//type AssetPutRequest struct {
//	Codigo string `json:"codigo"`
//	Nome   string `json:"nome"`
//	Logo   string `json:"logo"`
//}

type OrderPostRequest struct {
	Ativo      int64 	`json:"id_ativo"`
	Quantidade int 		`json:"quantidade"`
	Valor      float32	`json:"valor"`
}
