package boaconstructor

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	dlg "github.com/sqweek/dialog"
	data "github.com/westarver/boa-constructor/appdata"
	"github.com/westarver/boa-constructor/io"
	"github.com/westarver/boa-constructor/ui"
)

func Run() {
	data.GlobalData.MainWindow.CenterOnScreen()
	data.GlobalData.Icon = ui.ResourceBoaIconPng

	tabs := container.NewAppTabs(
		container.NewTabItem("App Info", ui.AppInfoTab()),
		container.NewTabItem("Commands/Flags", ui.CommandTab()),
		container.NewTabItem("More", ui.MoreTab()),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	data.GlobalData.MainWindow.SetContent(tabs)
	data.GlobalData.MainWindow.SetCloseIntercept(func() {
		verifyExit()
	})
	data.GlobalData.MainWindow.SetIcon(ui.ResourceBoaIconPng)
	data.GlobalData.MainWindow.SetPadded(false)
	ui.SetMainMenu()
	Runcli(os.Stderr)
	loadOnStart()
	if !data.GlobalData.MainList.Empty() {
		ui.RefreshPreview()
	}
	app := data.GlobalData.App
	ht := app.Preferences().FloatWithFallback("Height", 650)
	wd := app.Preferences().FloatWithFallback("Width", 1300)

	data.GlobalData.MainWindow.Resize(fyne.NewSize(float32(wd), float32(ht)))
	data.GlobalData.MainWindow.ShowAndRun()
}

func verifyExit() {
	app := data.GlobalData.App
	app.Preferences().SetFloat("Height", float64(data.GlobalData.MainWindow.Canvas().Size().Height))
	app.Preferences().SetFloat("Width", float64(data.GlobalData.MainWindow.Canvas().Size().Width))
	if data.GlobalData.Dirty {
		ok := dlg.Message("%s", "The latest edits have not been saved. Do you want to exit?").Title("Work not saved").YesNo()
		if ok {
			data.GlobalData.MainWindow.Close()
		}
	} else {
		data.GlobalData.MainWindow.Close()
	}
}

func loadOnStart() {
	file := data.GlobalData.LoadedPath
	if file != "" {
		err := io.DoLoadJSON(file)
		if err != nil {
			log.Println(err)
			os.Exit(999)
		}
		if head := data.GlobalData.MainList.Head(); head != nil {
			ui.PopulateTab(head.Name())
		}
	}
	data.GlobalData.Dirty = false
}
