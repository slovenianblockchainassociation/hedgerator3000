package hedger

import (
	"hedgerator3000/connector"
	"testing"
)

func TestHedge(t *testing.T) {
	walletPublicAddress := ""
	conn := connector.New(nil)

	err := Hedge(walletPublicAddress, conn)
	if err != nil {
		t.Error(err)
	}
}
