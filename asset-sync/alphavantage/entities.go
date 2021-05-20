package alphavantage

type GlobalQuote struct {
	GlobalQuoteData `json:"Global Quote"`
}

type GlobalQuoteData struct {
	Price	string	`json:"05. price"`
}