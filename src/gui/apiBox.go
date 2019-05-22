package gui

import (
	"github.com/andlabs/ui"
	"log"
	"ethe"
	"strings"
	"fmt"
)

var (
	apiIndexLine *ui.Spinbox
	apiPwdLine   *ui.Entry
	apiUrlLine   *ui.Entry
	apiMuliLine  *ui.MultilineEntry
)

type api struct {
	gethURL string
	box     *ui.Box
}

func newApiBox(url string) api {
	if url == "" {
		log.Fatal("URL is Empty")
	} else if !strings.HasPrefix(url, "http://") {
		log.Fatal("ERROR Prefix! ex> 'http://url:port'")
	}

	return api{
		gethURL: url,
		box: func() *ui.Box {
			v := ui.NewVerticalBox()
			v.SetPadded(true)
			return v
		}()}
}

func (a api) show() *ui.Box {
	log.Println("Setting api tab.")
	a.init()
	a.selectWallet()
	a.myInfo()
	return a.box
}

func (a api) init() {
	apiUrlLine = ui.NewEntry()
	apiUrlLine.SetText(a.gethURL)
	apiUrlLine.SetReadOnly(true)
	apiIndexLine = ui.NewSpinbox(0, 100)
	apiPwdLine = ui.NewPasswordEntry()
	apiMuliLine = ui.NewMultilineEntry()
	apiMuliLine.SetReadOnly(true)

}

func (a api) selectWallet() {
	h := ui.NewHorizontalBox()

	g := ui.NewGroup("")

	from := ui.NewForm()
	from.Append("gethURL :  ", apiUrlLine, false)
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

func requestGethAPI() {
	apiMuliLine.SetText("")

	key, err := ethe.ResultKeystore(apiIndexLine.Value(), apiPwdLine.Text())
	if err != nil {
		log.Println(err)
		return
	}
	client, err := newGethApi(apiUrlLine.Text(), key.Address())
	if err != nil {
		log.Println(err)
		return
	}

	balance, err := client.getBalance()
	if err != nil {
		log.Println(err)
		return
	}

	apiMuliLine.Append(fmt.Sprintf("Account : 0x%0x\n", key.Address()))
	apiMuliLine.Append(fmt.Sprintf("Balance : %s\n", balance))
}

func (a api) myInfo() {
	a.box.Append(apiMuliLine, true)
}
