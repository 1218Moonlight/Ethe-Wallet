package jsonRPC

import (
	"github.com/onrik/ethrpc"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"bytes"
	"fmt"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"strconv"
	"strings"
	"log"
)

type Client struct {
	dial    *ethrpc.EthRPC
	chainID *big.Int
}

func NewClient(url string, chainID *big.Int) Client {
	if url == "" {
		log.Fatal("URL is Empty")
	} else if !strings.HasPrefix(url, "http://") {
		log.Fatal("ERROR Prefix! ex> 'http://url:port'")
	}

	return Client{
		dial:    ethrpc.New(url),
		chainID: chainID,
	}
}

func (c Client) SendRawTransaction(bytePri []byte, nonce int, from string, to common.Address,
	amount string, gasLimit int, gasPrice int) (string, error) {

	priKey, err := crypto.HexToECDSA(hex.EncodeToString(bytePri))
	if err != nil {
		return "", err
	}

	a, err := strconv.ParseInt(amount, 10, 64)
	value := big.NewInt(a)
	gasprice := big.NewInt(int64(gasPrice))

	tx := types.NewTransaction(uint64(nonce), to, value, uint64(gasLimit), gasprice, nil)

	signer := types.NewEIP155Signer(c.chainID)
	signedTx, err := types.SignTx(tx, signer, priKey)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	signedTx.EncodeRLP(&buff)

	s, err := c.dial.EthSendRawTransaction(fmt.Sprintf("0x%x", buff.Bytes()))
	if err != nil {
		return "", err
	}

	return s, nil
}
