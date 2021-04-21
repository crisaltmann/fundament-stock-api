package domain

type Portfolio struct {
	Id int64
	UserId int64
	Nome string
}

type Quarter struct {
	Id int64
	Codigo string
}

type PortfolioItem struct {
	Id int64
	IdPortfolio int64
	IdAtivo int64
	Qtde int32
}

type PortfolioQuarterItem struct {
	Id int64
	IdQuater int64
	IdPortfolio int64
	IdAtivo int64
	Qtde int64
}

type Ordem struct {
	Id int64
	IdPorfolio int64
	IdAtivo int64
	Quantidade int64
	ValorCompra int64
}


