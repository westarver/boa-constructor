package boaconstructor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/westarver/boa"
)

var appDataName = boa.AppDataName() //"BOA-APP-DATA"

type packageGlobal struct {
	app        fyne.App
	mainWindow fyne.Window
	icon       *fyne.Resource
	preview    *readOnlyEntry
	appName    *widget.Entry
	pkg        *widget.Entry
	author     *widget.Entry
	loadedPath string
	dirty      bool
}

func newPackageGlobal() *packageGlobal {
	app := app.New()
	win := app.NewWindow("Boa Constructor")

	return &packageGlobal{
		app:        app,
		mainWindow: win,
		preview:    nil,
		appName:    nil,
		pkg:        nil,
		author:     nil,
		loadedPath: "",
		dirty:      false,
	}
}

var appData = newPackageGlobal()

func Run() {
	appData.mainWindow.CenterOnScreen()
	appData.mainWindow.Resize(fyne.Size{Width: 800, Height: 600})
	res, _ := fyne.LoadResourceFromPath("/home/westarver/go/src/boa-constructor-work/boa-constructor-cmd/assets/boa-icon.png")
	appData.icon = &res

	tabs := container.NewAppTabs(
		container.NewTabItem("App Info", appInfoTab()),
		container.NewTabItem("Commands", commandTab()),
		container.NewTabItem("Flags", flagTab()),
		container.NewTabItem("More", moreTab()),
	)

	tabs.SetTabLocation(container.TabLocationTop)
	appData.mainWindow.SetContent(tabs)
	appData.mainWindow.SetCloseIntercept(func() {
		verifyExit()
	})
	appData.mainWindow.SetIcon(res)
	setMainMenu()
	appData.mainWindow.ShowAndRun()
}

func verifyExit() {
	if appData.dirty {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to exit?", func(b bool) {
			if b {
				appData.mainWindow.Close()
			}
		}, appData.mainWindow)
		dlg.Show()
	} else {
		appData.mainWindow.Close()
	}
}
