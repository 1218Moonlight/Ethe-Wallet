package gui

import (
	"github.com/andlabs/ui"
	"io/ioutil"
	"ethe"
	"log"
)

var (
	apiIndexLine *ui.Spinbox
	apiPwdLine   *ui.Entry
	urlLine      *ui.Entry
)

type api struct {
	box *ui.Box
}

func newApiBox() api {
	return api{
		box: func() *ui.Box {
			v := ui.NewVerticalBox()
			v.SetPadded(true)
			return v
		}()}
}

func (a api) show() *ui.Box {
	a.init()
	a.selectWallet()
	return a.box
}

func (a api) init() {
	urlLine = ui.NewEntry()
	apiIndexLine = ui.NewSpinbox(0, 100)
	apiPwdLine = ui.NewPasswordEntry()

}

func (a api) selectWallet() {
	h := ui.NewHorizontalBox()

	g := ui.NewGroup("")

	from := ui.NewForm()
	from.Append("gethURL :  ", urlLine, false)
	from.Append("index :  ", apiIndexLine, true)
	from.Append("pwd :  ", apiPwdLine, true)
	g.SetChild(from)

	h.Append(g, true)

	apiBtn := ui.NewButton("request")
	apiBtn.OnClicked(func(button *ui.Button) {
		requestGethAPI()
	})

	h.Append(apiBtn, false)

	a.box.Append(h, false)
}

func requestGethAPI(){
	eoaListSearch(apiIndexLine.Value())
	//newGethApi(urlLine.Text(), )
}

func eoaListSearch(index int) {
	files, err := ioutil.ReadDir(ethe.WalletPath)
	if err != nil {
		log.Fatal(err)
	}

	if (index + 1) != len(files) {
		log.Println("index out of range")
		return
	}

	for i, file := range files {
		if i == index{
			log.Println(file.Name())
		}
	}
}