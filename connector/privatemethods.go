package connector

import "errors"

func (api *HttpEngine) SendOrder(symbol string, price, quantity float64, postOnly bool) (*Order, error) {
	return nil, errors.New("SendOrder() not implemented")
}

func (e *HttpEngine) Positions() ([]Position, error) {
	return nil, errors.New("Positions() not implemented")
}

func (E *HttpEngine) UserWallet() (*Wallet, error) {
	return nil, errors.New("UserWallet() not implemented")
}
