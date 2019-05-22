package gui

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"context"
	"math/big"
)

type gethApi struct {
	dial *ethclient.Client
	eoa  common.Address
}

func newGethApi(url string, eoa common.Address) gethApi {
	return gethApi{
		dial: gethApiClient(url),
		eoa:  eoa,
	}
}

func gethApiClient(url string) (*ethclient.Client) {
	cli, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

func (g gethApi) getBalance() (*big.Int, error) {
	defer g.dial.Close()

	ct := context.Background()

	balance, err := g.dial.BalanceAt(ct, g.eoa, nil)
	if err != nil {
		return big.NewInt(0), err
	}

	return balance, nil
}
