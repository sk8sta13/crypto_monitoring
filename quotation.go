package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"time"
	"context"
	"errors"
	"fmt"
	"strconv"
)

type Quotation struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func NewQuotation(url string) (*Quotation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500 * time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("A request to\"", url, "\"exceeded the 500 ms timeout.")
		}
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quote Quotation
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

func (q *Quotation) Alert(rule Rule) (bool, error) {
	price, err := strconv.ParseFloat(q.Price, 64)
	if err != nil {
		log.Println("Error converting string \"%s\".", rule.Value)
		return false, fmt.Errorf("Error converting string \"%s\".", rule.Value)
	}

	switch rule.Operator {
	case "<":
		return price < rule.Value, nil
	case "<=":
		return price <= rule.Value, nil
	case "=":
		return price == rule.Value, nil
	case ">=":
		return price >= rule.Value, nil
	case ">":
		return price > rule.Value, nil
	case "<>":
		return price != rule.Value, nil
	default:
		log.Println("the operator \"", rule.Operator, "\" setting is incorrect.")
		return false, fmt.Errorf("the operator \"%s\" setting is incorrect.", rule.Operator)
	}
}