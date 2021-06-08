package portfolio_domain

type Portfolio struct {
	Ativo      Ativo
	Quantidade int
	Valor      float32
	Usuario    int64
}

type Ativo struct {
	Id			int64
	Codigo		string
	Logo		string
	Total		int64
	Cotacao		float32
}