package asset_api

type AssetResponse struct {
	Id     	int64  `json:"id"`
	Codigo 	string `json:"codigo"`
	Nome   	string `json:"nome"`
	Logo   	string `json:"logo"`
	Cotacao float32 `json:"cotacao"`
}

type AssetPostRequest struct {
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`
	Logo   string `json:"logo""`
}

type AssetPutRequest struct {
	Codigo  string `json:"codigo"`
	Nome    string `json:"nome"`
	Logo    string `json:"logo"`
	Cotacao float32 `json:"cotacao"`
}
