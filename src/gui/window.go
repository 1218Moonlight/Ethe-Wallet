package gui

import (
	"github.com/andlabs/ui"
	"math/big"
)

const (
	textAccountTab = "Wallet"
	textApiTab     = "API"
	textTxTab      = "TX"
)

type window struct {
	main   *ui.Window
	tab    *ui.Tab
	eoaTab eoa
	apiTab api
	txTab  transaction
}

func newWindow(title string, width, height int, hashMenubar bool, gethURL string, chainID *big.Int) window {
	return window{
		main:   func() *ui.Window { return ui.NewWindow(title, width, height, hashMenubar) }(),
		tab:    func() *ui.Tab { return ui.NewTab() }(),
		eoaTab: newMainBox(),
		apiTab: newApiBox(gethURL),
		txTab:  newTxBox(gethURL, chainID)}
}

// Wallet Tab
func (w window) walletUI() {
	w.tab.Append(textAccountTab, w.eoaTab.show())
	w.tab.SetMargined(0, true)
}

func (w window) apiUI() {
	w.tab.Append(textApiTab, w.apiTab.show())
}

func (w window) txUI() {
	w.tab.Append(textTxTab, w.txTab.show())
}

func (w window) mainExit() {
	w.main.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	ui.OnShouldQuit(func() bool {
		w.main.Destroy()
		return true
	})
}

func (w window) mainShow() {
	w.main.SetChild(w.tab)
	w.main.Show()
}
