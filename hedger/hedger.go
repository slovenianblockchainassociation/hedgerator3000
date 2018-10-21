package hedger

import (
	"errors"
	"fmt"
	"hedgerator3000/connector"
	"hedgerator3000/wallet"
	"math"
)

const (
	contract = "XBTZ18"
	minDiff  = 0.001
	sell     = "Sell"
	buy      = "Buy"
)

func Hedge(walletPublicAddress string, conn *connector.HttpEngine) error {
	// get btc wallet balance
	currentWalletBalance, err := wallet.GetWalletBalance(walletPublicAddress)
	if err != nil {
		return errors.New(fmt.Sprintf("error while getting current wallet position: %s", err))
	}

	// get current hedge position
	// bitmex uses inverted btc contract, which means that one XBTZ18 contract is worth $1
	// position below is quoted in BTC, which you get by multiplying number of contracts with (average) entry price
	// negative number means we have a short position
	currentBtcHedgePosition, err := conn.GetCurrentSymbolPosition(contract)
	if err != nil {
		return errors.New(fmt.Sprintf("error while getting current hedge position: %s", err))
	}

	// TODO add bitmex balance check. decide how much cash should be dedicated for hedge

	// if unhedged position is positive, it means we need to add more shorts on futures to get a neutral position
	unhedgedBtc := currentWalletBalance + currentBtcHedgePosition

	// if unhedged position is smaller then minDiff exit program
	if math.Abs(unhedgedBtc) < minDiff {
		return errors.New("exiting hedge: your position is hedged")
	}

	// to calculate exact number of contracts to hedge (since they are quoted in $1) we need to get trading prices
	orderBook, err := conn.OrderBookL2(contract, 10)
	if err != nil {
		return errors.New(fmt.Sprintf("error while getting order book: %s", err))
	}

	// get limit order trade price and number of contracts
	var tradePrice, tradeContractAmount, tradeBtcAmount float64

	for _, obLevel := range orderBook {
		// bitmex l2 orderbook is a slice of order book levels (asks and bids) sorted by price. first skip all asks
		if obLevel.Side == sell {
			continue
		}
		if obLevel.Size/obLevel.Price+tradeBtcAmount > unhedgedBtc {
			tradePrice = obLevel.Price
			tradeContractAmount += obLevel.Size
			break
		}
		tradeContractAmount += obLevel.Size
		tradeBtcAmount += obLevel.Size / obLevel.Price
	}

	// execute trade
	_, err = conn.SendOrder(contract, tradePrice, tradeContractAmount, false)
	if err != nil {
		return errors.New(fmt.Sprintf("error while sending hedge order (%f contracts @ %f): %s", tradeContractAmount, tradePrice, err))
	}

	return nil
}
