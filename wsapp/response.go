package wsapp

import (
	"aybjax_ascendex_websocket/app"
	"strconv"
)

// {"m":"bbo","symbol":"BTC/USDT","data":{"ts":1681765877548,"bid":["29412.54","0.01346"],"ask":["29459.06","0.01746"]}}

type response struct {
	Type string       `json:"m"`
	Data responseData `json:"data"`
}

type responseData struct {
	Bid [2]string `json:"bid"` //price, amount
	Ask [2]string `json:"ask"` //price, amount
}

func (r *response) isTypeBBO() bool {
	return r.Type != "" && r.Type == "bbo"
}

func (r *response) getBestOrderBook() (app.BestOrderBook, error) {
	var result app.BestOrderBook
	var err error

	result.Ask.Price, err = strconv.ParseFloat(r.Data.Ask[0], 64)

	if err != nil {
		return result, err
	}

	result.Ask.Amount, err = strconv.ParseFloat(r.Data.Ask[1], 64)

	if err != nil {
		return result, err
	}

	result.Bid.Price, err = strconv.ParseFloat(r.Data.Bid[0], 64)

	if err != nil {
		return result, err
	}

	result.Bid.Amount, err = strconv.ParseFloat(r.Data.Bid[1], 64)

	if err != nil {
		return result, err
	}

	return result, nil
}
