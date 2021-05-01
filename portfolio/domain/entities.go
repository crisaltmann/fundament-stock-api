package orderdomain

type Quarter struct {
	Id int64
	Codigo string
}

type PortfolioItem struct {
	Id int64
	IdUser int64
	IdAtivo int64
	Qtde int32
}

type PortfolioQuarterItem struct {
	Id int64
	IdQuater int64
	IdUser int64
	IdAtivo int64
	Qtde int64
}

type Order struct {
	Id int64
	IdPorfolio int64
	Data int64
	IdAtivo int64
	Quantidade int64
	Valor int64
}


