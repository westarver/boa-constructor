package ui

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/widget"
	data "github.com/westarver/boa-constructor/appdata"
	"github.com/westarver/fynewidgets"
)

// appInfoTab is the code to create and populate the first Tab
// where application data such as app name and package name is
// gathered.  The preview screen is presented on this tab as well
func AppInfoTab() *fyne.Container {
	//buttonInfoLabel.SetText("")
	toolbar := widget.NewToolbar(
		newToolbarAction("Refresh Preview", ResourcePreview40dpSvg, RefreshPreview),
		newToolbarAction("File Open      ", ResourceFileopen40dpSvg, doFileOpen),
		newToolbarAction("File Save      ", ResourceSave40dpSvg, doFileSave),
		newToolbarAction("Save As        ", ResourceSaveas40dpSvg, doFileSaveAs),
	)

	snake := widget.NewButtonWithIcon("", data.GlobalData.Icon, doHelpAbout)
	//infoLabel1 := widget.NewLabelWithStyle("Application Data", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	infoBox1 := container.NewHBox(toolbar /*buttonInfoLabel*/, layout.NewSpacer(), snake)
	//infoBox2 := container.NewHBox(layout.NewSpacer(), infoLabel1, layout.NewSpacer())
	info1 := widget.NewLabel("Application Name")
	//info1.TextStyle = fyne.TextStyle{}
	text1 := widget.NewEntry()
	text1.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	info2 := widget.NewLabel("Package Name For Generated Help Code")
	//info2.TextStyle = fyne.TextStyle{Bold: true}
	text2 := widget.NewEntry()
	text2.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	info25 := widget.NewLabel("Package Name For Generated Run Code")
	//info25.TextStyle = fyne.TextStyle{Bold: true}
	text25 := widget.NewEntry()
	text25.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	info26 := widget.NewLabel("Module Import Path(top of the go.mod file)")
	//info26.TextStyle = fyne.TextStyle{Bold: true}
	text26 := widget.NewEntry()
	text26.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	info3 := widget.NewLabel("Author")
	//info3.TextStyle = fyne.TextStyle{Bold: true}
	text3 := widget.NewEntry()
	text3.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	f := container.New(layout.NewFormLayout(), info1, text1, info2, text2, info25, text25, info26, text26, info3, text3)
	form1 := container.NewVBox(infoBox1, f)

	divider1 := widget.NewSeparator()
	divider2 := widget.NewSeparator()

	infoLabel2 := widget.NewLabelWithStyle("Preview (read only)", fyne.TextAlignCenter, fyne.TextStyle{Bold: false})
	text4 := fynewidgets.NewReadOnlyEntry()
	text4.SetText(strings.Repeat("\n", previewSize))
	text4.SetMinRowsVisible(10)
	preview := container.NewScroll(container.NewBorder(nil, nil, nil, nil, text4))
	preview.SetMinSize(fyne.Size{
		Width:  800,
		Height: 260,
	})

	data.GlobalData.AppName = text1
	data.GlobalData.HelpPkg = text2
	data.GlobalData.RunPkg = text25
	data.GlobalData.ImportPath = text26
	data.GlobalData.Author = text3
	data.GlobalData.Preview = text4

	return container.NewVBox(form1, divider1, infoLabel2, preview, layout.NewSpacer(), divider2)
}

func ResetAll() {
	data.GlobalData.MainList.ResetAll()
	data.GlobalData.More.SetText("")
	data.GlobalData.AppName.SetText("")
	data.GlobalData.HelpPkg.SetText("")
	data.GlobalData.RunPkg.SetText("")
	data.GlobalData.Author.SetText("")
	data.GlobalData.Preview.SetText(strings.Repeat("\n", previewSize))
	data.GlobalData.LoadedPath = ""
	data.GlobalData.Dirty = false
	ClearTab()
}
