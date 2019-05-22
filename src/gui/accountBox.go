package gui

import (
	"github.com/andlabs/ui"
	"ethe"
	"io/ioutil"
	"fmt"
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

func (m eoa) ui() *ui.Box {
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

func (m eoa) setGenerateWallet() *ui.Box {
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
	return m.box
}

func (m eoa) accountReloadBTN() *ui.Box {
	reloadBtn = ui.NewButton(textReload)
	reloadBtn.OnClicked(func(button *ui.Button) {
		eoaListAppend()
	})

	m.box.Append(reloadBtn, false)
	return m.box
}

func (m eoa) setAccountList() *ui.Box {
	eoaList = ui.NewMultilineEntry()
	eoaList.SetReadOnly(true)

	eoaListAppend()

	m.box.Append(eoaList, true)
	return m.box
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
