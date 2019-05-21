package gui

import (
	"github.com/andlabs/ui"
	"ethe"
)

const (
	textNewAccount = "NewAccount"
)

var (
	pwdEntry  *ui.Entry
	newEoaBtn *ui.Button
)

type mainBox struct {
	top *ui.Box
}

func newMainBox() mainBox {
	box := mainBox{
		top: func() *ui.Box {
			b := ui.NewVerticalBox()
			b.SetPadded(true)
			return b
		}()}

	return box
}

func (m mainBox) setPwdText() *ui.Box {
	pwdEntry = ui.NewEntry()

	m.top.Append(pwdEntry, false)
	return m.top
}

func (m mainBox) setNewAccountBTN() *ui.Box {
	newEoaBtn = ui.NewButton(textNewAccount)
	newEoaBtn.OnClicked(func(button *ui.Button) {
		ethe.NewAccount(pwdEntry.Text())
	})

	m.top.Append(newEoaBtn, false)
	return m.top
}

func (m mainBox) ui() *ui.Box {
	m.setPwdText()
	m.setNewAccountBTN()
	return m.top
}
