package ethe

import (
	k "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
)

const walletPath string = "./wallet"

func NewAccount(pwd string) (accounts.Account, error) {
	a := accounts.Account{}
	ks := k.NewKeyStore(walletPath, 2, 1) // Todo: scryptN, scryptP
	a, err := ks.NewAccount(pwd)
	if err != nil {
		return a, err
	}
	return a, nil
}
