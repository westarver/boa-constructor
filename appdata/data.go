package appdata

import (
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/westarver/boa-constructor/clargs"
	"github.com/westarver/fynewidgets"
)

// unexported struct to hold references to widgets and other data
// needing to be accessed from all over the app.
type globalData struct {
	App        fyne.App
	MainWindow fyne.Window
	Icon       *fyne.StaticResource
	Preview    *fynewidgets.ReadOnlyEntry
	AppName    *widget.Entry
	HelpPkg    *widget.Entry
	RunPkg     *widget.Entry
	ImportPath *widget.Entry
	Author     *widget.Entry
	More       *widget.Entry
	LoadedPath string
	Dirty      bool
	IsGui      bool
	MdHelp     bool
	MainList   *clargs.ArgList
}

func newGlobalData() *globalData {
	app := app.NewWithID("com.github.boa.jimmy")
	win := app.NewWindow("Boa Constructor")

	return &globalData{
		App:        app,
		MainWindow: win,
		Preview:    nil,
		AppName:    nil,
		HelpPkg:    nil,
		RunPkg:     nil,
		ImportPath: nil,
		Author:     nil,
		More:       nil,
		LoadedPath: "",
		Dirty:      false,
		IsGui:      false,
		MdHelp:     false,
		MainList:   clargs.NewArgList(),
	}
}

var GlobalData = newGlobalData()

func NormalizeImportPath() string {
	importpath := GlobalData.ImportPath.Text
	return strings.TrimSuffix(importpath, "/")
}

func TruncateImportPath() string {
	return filepath.Dir(NormalizeImportPath())
}
