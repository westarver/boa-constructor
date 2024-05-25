package boaconstructor

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// appInfoTab is the code to create and populate the first Tab
// where application data such as app name and package name is
// gathered.  The preview screen is presented on this tab as well
func appInfoTab() *fyne.Container {
	snake := widget.NewButtonWithIcon("", *appData.icon, nil)
	infoLabel1 := widget.NewLabelWithStyle("Application Data", fyne.TextAlignTrailing, fyne.TextStyle{Bold: true})
	infoBox1 := container.NewHBox(layout.NewSpacer(), infoLabel1, layout.NewSpacer(), snake)
	info1 := widget.NewLabel("Application Name")
	info1.TextStyle = fyne.TextStyle{Bold: true}
	text1 := widget.NewEntry()
	text1.OnChanged = func(s string) { appData.dirty = true }

	info2 := widget.NewLabel("Package Name For Generated Code")
	info2.TextStyle = fyne.TextStyle{Bold: true}
	text2 := widget.NewEntry()
	text2.OnChanged = func(s string) { appData.dirty = true }
	info3 := widget.NewLabel("Author")
	info3.TextStyle = fyne.TextStyle{Bold: true}
	text3 := widget.NewEntry()
	text3.OnChanged = func(s string) { appData.dirty = true }
	f := container.New(layout.NewFormLayout(), info1, text1, info2, text2, info3, text3)
	form1 := container.NewVBox(infoBox1, f)

	divider1 := widget.NewSeparator()
	divider2 := widget.NewSeparator()

	infoLabel2 := widget.NewLabelWithStyle("Preview (read only)", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	text4 := newreadOnlyEntry()
	text4.SetText(strings.Repeat("\n", previewSize))
	preview := container.NewVScroll(container.NewVBox(text4, layout.NewSpacer()))
	preview.SetMinSize(fyne.Size{
		Width:  text4.Size().Width,
		Height: 400,
	})

	appData.appName = text1
	appData.pkg = text2
	appData.author = text3
	appData.preview = text4

	button3 := widget.NewButton("Generate Source Files", func() {
		confirmGen()
	})
	button4 := widget.NewButton("Open", func() {
		loadFromFile(true)
	})
	button5 := widget.NewButton("Save", func() {
		saveToFile(true)
	})
	btnbox := container.NewHBox(button4, button5, button3, layout.NewSpacer())
	return container.NewVBox(form1, divider1, infoLabel2, preview, layout.NewSpacer(), divider2, btnbox)
}

func resetAll() {
	cmdInfo.resetAll()
	flagInfo.resetAll()
	moreInfo.moreSec.SetText("")
	appData.appName.SetText("")
	appData.pkg.SetText("")
	appData.author.SetText("")
	appData.preview.SetText(strings.Repeat("\n", previewSize))
	appData.loadedPath = ""
	appData.dirty = false
}
