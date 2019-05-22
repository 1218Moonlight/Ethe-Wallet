package gui

import (
	"github.com/andlabs/ui"

	"ethe"
	"io/ioutil"
	"fmt"
	"log"
)

const (
	textNewAccount  = "newAccount"
	textAccountList = "Account List"
	textReload      = "reload"
)

var (
	pwdEntry  *ui.Entry
	newEoaBtn *ui.Button
	reloadBtn *ui.Button
	eoaList   *ui.MultilineEntry
)

type eoa struct {
	box *ui.Box
}

func newMainBox() eoa {
	box := eoa{
		box: func() *ui.Box {
			b := ui.NewVerticalBox()
			b.SetPadded(true)
			return b
		}()}

	return box
}

func (m eoa) show() *ui.Box {
	log.Println("Setting Wallet Tab")
	m.setGenerateWallet()
	m.hSeparator(textAccountList)
	m.accountReloadBTN()
	m.setAccountList()
	return m.box
}

func (m eoa) hSeparator(label string) {
	m.box.Append(ui.NewHorizontalSeparator(), false)
	m.box.Append(ui.NewLabel(label), false)
}

func (m eoa) setGenerateWallet() {
	group := ui.NewGroup("EOA")
	group.SetMargined(true)

	h := ui.NewHorizontalBox()

	from := ui.NewForm()

	pwdEntry = ui.NewEntry()
	from.Append("PWD :  ", pwdEntry, true)

	h.Append(from, true)
	h.Append(setNewAccountBTN(), false)

	group.SetChild(h)

	m.box.Append(group, false)
}

func (m eoa) accountReloadBTN() {
	reloadBtn = ui.NewButton(textReload)
	reloadBtn.OnClicked(func(button *ui.Button) {
		eoaListAppend()
	})

	m.box.Append(reloadBtn, false)
}

func (m eoa) setAccountList() {
	eoaList = ui.NewMultilineEntry()
	eoaList.SetReadOnly(true)

	eoaListAppend()

	m.box.Append(eoaList, true)

}

func setNewAccountBTN() *ui.Button {
	newEoaBtn = ui.NewButton(textNewAccount)
	newEoaBtn.OnClicked(func(button *ui.Button) {
		ethe.NewAccount(pwdEntry.Text())
	})
	return newEoaBtn
}

func eoaListAppend() {
	eoaList.SetText("")

	files, err := ioutil.ReadDir(ethe.WalletPath)
	if err != nil {
		eoaList.Append(err.Error())
	}

	for i, file := range files {
		eoaList.Append(fmt.Sprintf("%d : %s\n", i, file.Name()))
	}
}
