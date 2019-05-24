package gui

import (
	"github.com/andlabs/ui"

	"math/big"
	"jsonRPC"
	"ethe"
	"log"
	"github.com/ethereum/go-ethereum/common"
)

var (
	txToEntry       *ui.Entry
	txValueEntry    *ui.Entry
	txGasLimitEntry *ui.Spinbox
	txGasPriceEntry *ui.Spinbox
	txFromEntry     *ui.Spinbox
	txNonceEntry    *ui.Spinbox
	txPwdEntry      *ui.Entry
	txSendButton    *ui.Button
)

type transaction struct {
	box *ui.Box
	rpc jsonRPC.Client
}

func newTxBox(url string, chainID *big.Int) transaction {
	return transaction{
		box: func() *ui.Box {
			v := ui.NewVerticalBox()
			v.SetPadded(true)
			return v
		}(),
		rpc: jsonRPC.NewClient(url, chainID),
	}
}

func (t transaction) init() {
	txToEntry = ui.NewEntry()
	txValueEntry = ui.NewEntry()
	txGasLimitEntry = ui.NewSpinbox(0, 999999)
	txGasPriceEntry = ui.NewSpinbox(0, 999999)
	txFromEntry = ui.NewSpinbox(0, 100)
	txPwdEntry = ui.NewPasswordEntry()
	txNonceEntry = ui.NewSpinbox(0, 999999)
	txSendButton = ui.NewButton("Send")
}

func (t transaction) show() *ui.Box {
	t.init()
	t.sendTx()
	return t.box
}

func (t transaction) sendTx() {
	g := ui.NewGroup("SendTransaction")

	f := ui.NewForm()
	f.Append("Nonce :  ", txNonceEntry, false)
	f.Append("From :  ", txFromEntry, false)
	f.Append("To :  ", txToEntry, false)
	f.Append("Value :  ", txValueEntry, false)
	f.Append("GasLimit :  ", txGasLimitEntry, false)
	f.Append("GasPrice :  ", txGasPriceEntry, false)
	f.Append("PWD :  ", txPwdEntry, false)
	f.Append("", txSendButton, false)

	t.clickTxSendButton()

	g.SetChild(f)

	t.box.Append(g, false)
}

func (t transaction) clickTxSendButton() {
	txSendButton.OnClicked(func(button *ui.Button) {
		key, _ := ethe.ResultKeystore(txFromEntry.Value(), txPwdEntry.Text())

		bytePri, err := key.PrivateKey()
		if err != nil {
			log.Println(err)
			return
		}

		to := common.HexToAddress(txToEntry.Text())

		txAccount, err := t.rpc.SendRawTransaction(bytePri, txNonceEntry.Value(), key.Address().String(), to, txValueEntry.Text(),
			txGasLimitEntry.Value(), txGasPriceEntry.Value())

		if err != nil {
			log.Println(err)
			return
		}

		log.Println(txAccount)
	})
}
