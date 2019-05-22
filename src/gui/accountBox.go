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

func (m eoa) ui() *ui.Box {
	m.setPwdText()
	m.setNewAccountBTN()
	return m.box
}
