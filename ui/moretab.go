package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	data "github.com/westarver/boa-constructor/appdata"
)

func MoreTab() *fyne.Container {
	infoLabel1 := widget.NewLabelWithStyle("More Help Text (freeform and optional)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	info1 := widget.NewLabelWithStyle("More:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	text1 := widget.NewMultiLineEntry()
	text1.SetMinRowsVisible(12)
	text1.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	data.GlobalData.More = text1
	f := container.NewVBox(info1, text1)
	form1 := container.NewVBox(infoLabel1, f)

	divider := widget.NewSeparator()

	button1 := widget.NewButton("Clear Text", func() { text1.SetText("") })
	btnbox := container.NewHBox(button1, layout.NewSpacer())

	return container.NewVBox(form1, layout.NewSpacer(), divider, btnbox)
}
