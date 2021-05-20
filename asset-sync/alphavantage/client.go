package alphavantage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

import (
	"fmt"
)

type Client struct {

}

func (c Client) GetGlobalQuote(stock string) (GlobalQuote, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=MGZL2MQARYZO1I67", stock+".SA")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GlobalQuote{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return GlobalQuote{}, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	quote := GlobalQuote{}
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return GlobalQuote{}, err
	}
	return quote, nil
}