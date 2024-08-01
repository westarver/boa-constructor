package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	data "github.com/westarver/boa-constructor/appdata"
)

var buttonInfoLabel = widget.NewLabel("")

type hoverButton struct {
	widget.Button
	popup *widget.PopUp
	msg   string
	done  chan bool
}

func newHoverButton(msg string, icon fyne.Resource, fn func()) *hoverButton {
	b := &hoverButton{}
	b.Importance = widget.LowImportance
	b.Icon = icon
	b.OnTapped = fn
	b.msg = msg
	b.popup = &widget.PopUp{Content: buttonInfoLabel,
		Canvas: data.GlobalData.MainWindow.Canvas(),
	}
	b.popup.Content.(*widget.Label).Text = b.msg
	return b
}

func (b *hoverButton) doPopup(done chan bool) {
	X := b.Position().X
	Y := b.Position().Y + 26

	b.popup.Content.(*widget.Label).Text = b.msg
	b.popup.ShowAtPosition(fyne.NewPos(X, Y))
	<-done
	b.popup.Hide()
}

func (b *hoverButton) MouseIn(e *desktop.MouseEvent) {
	time.Sleep(time.Millisecond * 30)
	b.done = make(chan bool)
	go b.doPopup(b.done)
}

func (b *hoverButton) MouseMove(e *desktop.MouseEvent) {
}

func (b *hoverButton) MouseOut() {
	defer close(b.done)
	b.done <- true
}
