package gui

import (
	"github.com/andlabs/ui"
	"ethe"
	"io/ioutil"
	"fmt"
)

const (
	textNewAccount  = "NewAccount"
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
	m.setPwdText()
	m.setNewAccountBTN()
	m.hSeparator(textAccountList)
	m.accountReloadBTN()
	m.setAccountList()
	return m.box
}

func (m eoa) hSeparator(label string) {
	m.box.Append(ui.NewHorizontalSeparator(), false)
	m.box.Append(ui.NewLabel(label), false)
}

func (m eoa) setPwdText() *ui.Box {
	pwdEntry = ui.NewEntry()

	m.box.Append(pwdEntry, false)
	return m.box
}

func (m eoa) setNewAccountBTN() *ui.Box {
	newEoaBtn = ui.NewButton(textNewAccount)
	newEoaBtn.OnClicked(func(button *ui.Button) {
		ethe.NewAccount(pwdEntry.Text())
	})

	m.box.Append(newEoaBtn, false)
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
