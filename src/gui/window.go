package gui

import (
	"github.com/andlabs/ui"
)

const (
	textAccountTab = "Wallet"
)

type window struct {
	main   *ui.Window
	tab    *ui.Tab
	eoaTab eoa
}

func newWindow(title string, width, height int, hashMenubar bool) window {
	return window{
		main:   func() *ui.Window { return ui.NewWindow(title, width, height, hashMenubar) }(),
		tab:    func() *ui.Tab { return ui.NewTab() }(),
		eoaTab: newMainBox()}
}

// Wallet Tab
func (w window) accountTab() {
	w.tab.Append(textAccountTab, w.eoaTab.ui())
	w.tab.SetMargined(0, true)
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