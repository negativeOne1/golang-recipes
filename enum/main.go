package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type currencyStruct struct {
	Cur Currency `json:"currency"`
}

func main() {
	fmt.Println(EUR)

	j := []byte(`{"currency": "USD"}`)
	var c currencyStruct
	if err := json.Unmarshal(j, &c); err != nil {
		panic(err)
	}

	fmt.Println(c.Cur)

}

type Currency int

const (
	EUR Currency = iota
	USD
	CHF
)

var CurrencyStrings []string = []string{"EUR", "USD", "CHF"}

func (c Currency) String() string {
	return CurrencyStrings[c]
}

func (c Currency) MarshalJSON() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *Currency) UnmarshalJSON(b []byte) (err error) {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	*c, err = CurrencyFromText(s)
	return err
}

func CurrencyFromText(s string) (Currency, error) {
	s = strings.ToUpper(s)
	if len(s) != 3 {
		return 0, errors.New("too long")
	}

	for i, v := range CurrencyStrings {
		if v == s {
			return Currency(i), nil
		}
	}
	return 0, errors.New("not Known")
}
