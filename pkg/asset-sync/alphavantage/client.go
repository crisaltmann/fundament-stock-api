package alphavantage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

import (
	"fmt"
)

type Client struct {}

func NewAlphaVantageClient() Client {
	return Client{}
}

func (c Client) GetStockPrice(stock string) (float32, time.Time, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=MGZL2MQARYZO1I67", stock+".SA")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0.0, time.Time{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0.0, time.Time{}, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	quote := GlobalQuote{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return 0.0, time.Time{}, err
	}
	price, err := strconv.ParseFloat(quote.Price, 32)
	return float32(price), quote.LastTradeDay.ToTime(), err
}