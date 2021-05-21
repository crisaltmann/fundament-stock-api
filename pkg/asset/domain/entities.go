package asset_domain

import "time"

type Asset struct {
	Id     	int64
	Codigo 	string
	Nome   	string
	Logo   	string
	Cotacao float32
}

type AssetPrice struct {
	Id 		int64
	Ativo   int64
	Cotacao float32
	Data    time.Time
}
