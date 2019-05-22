package gui

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"config"
	"log"
)

func Main() {
	log.Println("Start the Ethe-Wallet.")
	ui.Main(setupUI)
}

func setupUI() {

	log.Println("Reading the setup file.")
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	win := newWindow(c.Title, c.Width, c.Height, c.HashMenubar, c.Geth)

	win.mainExit()

	win.walletUI()

	win.apiUI()

	win.mainShow()
}
