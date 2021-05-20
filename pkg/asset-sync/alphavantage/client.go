package alphavantage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

import (
	"fmt"
)

type Client struct {}

func NewAlphaVantageClient() Client {
	return Client{}
}

func (c Client) GetStockPrice(stock string) (float32, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=MGZL2MQARYZO1I67", stock+".SA")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0.0, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0.0, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	quote := GlobalQuote{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return 0.0, err
	}
	price, err := strconv.ParseFloat(quote.Price, 32)
	return float32(price), err
}