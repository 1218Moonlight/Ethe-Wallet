package gui

import (
	"github.com/andlabs/ui"
)

type window struct {
	main     *ui.Window
	mainBox mainBox
}

func newWindow(title string, width, height int, hashMenubar bool) window {
	return window{
		main: func() *ui.Window { return ui.NewWindow(title, width, height, hashMenubar) }(),
		mainBox: newMainBox()}
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
	w.main.Show()
}

func (w window) setMainBox() {
	w.main.SetChild(w.mainBox.ui())
}