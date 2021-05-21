package alphavantage

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type JSONTime time.Time

type GlobalQuote struct {
	GlobalQuoteData `json:"Global Quote"`
}

type GlobalQuoteData struct {
	Codigo           string 	`json:"01. symbol"`
	Price	         string		`json:"05. price"`
	LastTradeDay     JSONTime 	`json:"07. latest trading day"`
}

const DefaultFormat = time.RFC3339

var layouts = []string{
	DefaultFormat,
	"2006-01-02",
}

func (jt *JSONTime) String() string {
	t := time.Time(*jt)
	return t.Format(DefaultFormat)
}

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`,jt.String())), nil
}

func (jt *JSONTime) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), `"`)
	for _, layout := range layouts {
		t, err := time.Parse(layout, timeString)
		if err == nil {
			*jt = JSONTime(t)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Invalid date format: %s", timeString))
}

func (jt *JSONTime) ToTime() time.Time {
	return time.Time(*jt)
}