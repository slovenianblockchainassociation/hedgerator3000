package connector

import "errors"

func (api *HttpEngine) OrderBookL2(symbol string, depth int) ([]OrderBookL2, error) {
	return nil, errors.New("OrderBookL2() not implemented")
}
