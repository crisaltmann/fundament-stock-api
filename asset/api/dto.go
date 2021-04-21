package api

type AssetResponse struct {
	Id int64 `json:"id"`
	Codigo string `json:"codigo"`
	Nome string `json:"nome"`
}

type AssetPostRequest struct {
	Codigo string `json:"codigo"`
	Nome string `json:"nome"`
}

type AssetPutRequest struct {
	Codigo string `json:"codigo"`
	Nome string `json:"nome"`
}
