package portfolio_api

type PortfolioGetResponse struct {
	Ativo      Ativo		`json:"ativo"`
	Quantidade int			`json:"quantidade"`
	Valor      float32		`json:"valor"`
	Usuario	   int64		`json:"id_usuario"`
}

type Ativo struct {
	Id			int64		`json:"id"`
	Codigo		string		`json:"codigo"`
	Logo		string		`json:"logo"`
}
