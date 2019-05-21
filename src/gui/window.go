package gui

import (
	"config"
	"log"
	"github.com/andlabs/ui"
)

func setupUI() {
	a, e := config.File("config.json")
	if e !=nil {
		log.Fatal(e)
	}

	mainwin := ui.NewWindow(a.Title, a.Width, a.Height, a.HashMenubar)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	mainwin.Show()
}
