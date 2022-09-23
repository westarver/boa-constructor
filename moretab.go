package boaconstructor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type moreInfoStore struct {
	moreSec *widget.Entry
}

func moreTab() *fyne.Container {
	infoLabel1 := widget.NewLabelWithStyle("More Help Text (freeform and optional)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	info1 := widget.NewLabelWithStyle("More:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	text1 := widget.NewMultiLineEntry()
	text1.SetMinRowsVisible(12)
	text1.OnChanged = func(s string) { appData.dirty = true }
	moreInfo.moreSec = text1
	f := container.NewVBox(info1, text1)
	form1 := container.NewVBox(infoLabel1, f)

	divider := widget.NewSeparator()

	button1 := widget.NewButton("Clear Text", func() { text1.SetText(""); text1.Refresh() })
	button2 := widget.NewButton("Save", func() { saveToFile(true) })
	btnbox := container.NewHBox(button1, button2, layout.NewSpacer())

	return container.NewVBox(form1, layout.NewSpacer(), divider, btnbox)
}
