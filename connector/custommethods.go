package connector

import "errors"

func (api *HttpEngine) GetCurrentSymbolPosition(symbol string) (float64, error) {
	return 0, errors.New("GetCurrentSymbolPosition() not implemented")
}
