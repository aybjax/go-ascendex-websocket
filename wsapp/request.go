package wsapp

import (
	"fmt"
	"strings"
)

type request struct {
	Op string `json:"op"`
	Ch string `json:"ch,omitempty"`
}

func (request) getBBOSubMessage(symbol string) (*request, error) {
	vals := strings.Split(symbol, "_")

	if len(vals) != 2 || vals[0] == "" || vals[1] == "" {
		return nil, fmt.Errorf("Given symbol (%s) not in correct format", symbol)
	}

	ch := "bbo:" + strings.Join([]string{vals[1], vals[0]}, "/")

	return &request{
		Op: "sub",
		Ch: ch,
	}, nil
}

func (request) getPingMessage() *request {
	return &request{
		Op: "ping",
	}
}
