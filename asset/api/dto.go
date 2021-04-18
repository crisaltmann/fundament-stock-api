package api

type Asset struct {
	Id int64 `json:"id"`
	Codigo string `json:"codigo"`
	Nome string `json:"nome"`
}
