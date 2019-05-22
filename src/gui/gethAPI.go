package gui

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"context"
	"math/big"
)

type gethApi struct {
	dial *ethclient.Client
	eoa  common.Address
}

func newGethApi(url string, eoa common.Address) (gethApi, error) {
	client, err := gethApiClient(url)
	if err != nil {
		return gethApi{}, err
	}

	return gethApi{
		dial: client,
		eoa:  eoa,
	}, nil
}

func gethApiClient(url string) (*ethclient.Client, error) {
	cli, err := ethclient.Dial(url)
	if err != nil {
		return &ethclient.Client{}, err
	}
	return cli, nil
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
