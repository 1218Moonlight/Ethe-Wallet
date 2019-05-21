package gui

import (
	"github.com/andlabs/ui"
)

type window struct {
	main *ui.Window
}

func NewWindow(title string, width, height int, hashMenubar bool) window {
	return window{main: func() *ui.Window { return ui.NewWindow(title, width, height, hashMenubar) }()}
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
