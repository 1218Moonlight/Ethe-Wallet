package gui

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"config"
	"log"
)

func Main() {
	ui.Main(setupUI)
}

func setupUI() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	win := newWindow(c.Title, c.Width, c.Height, c.HashMenubar)

	win.mainExit()

	win.setMainBox()

	win.mainShow()
}
