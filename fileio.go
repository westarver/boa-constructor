package boaconstructor

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/westarver/boa"
)

type listable struct {
	fyne.URI
}

func (l listable) List() ([]fyne.URI, error) {
	return storage.List(l.URI)
}

func saveToFile(isJSON bool) {
	var err error
	path := appData.loadedPath

	if len(path) == 0 {
		path, err = os.Getwd()
		if err != nil {
			dialog.NewError(err, appData.mainWindow).Show()
			return
		}
		path += "~"
	}

	uri, err := storage.Parent(storage.NewFileURI(path))
	if err != nil {
		dialog.NewError(err, appData.mainWindow).Show()
		return
	}
	luri := listable{URI: uri}
	var fdialog *dialog.FileDialog

	if isJSON {
		fdialog = dialog.NewFileSave(doSaveJSON, appData.mainWindow)
		fdialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	} else {
		fdialog = dialog.NewFileSave(doSave, appData.mainWindow)
		fdialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".help"}))
	}

	fdialog.SetLocation(luri)
	fdialog.SetFileName(filepath.Base(appData.loadedPath))

	fdialog.Show()
}

func doSave(uwc fyne.URIWriteCloser, e error) {
	if uwc == nil || e != nil {
		return
	}
	allText := previewString()
	_, err := uwc.Write([]byte(allText))
	if err == nil {
		appData.dirty = false
	}
	uwc.Close()
}

func doSaveJSON(uwc fyne.URIWriteCloser, e error) {
	if uwc == nil || e != nil {
		return
	}

	jsonBytes := slice2json()
	_, err := uwc.Write(jsonBytes)
	if err == nil {
		appData.dirty = false
	}
	uwc.Close()
}

func loadFromFile(isJSON bool) {
	var err error
	path := appData.loadedPath

	if len(path) == 0 {
		path, err = os.Getwd()
		if err != nil {
			dialog.NewError(err, appData.mainWindow).Show()
			return
		}
		path += "~"
	}

	uri, err := storage.Parent(storage.NewFileURI(path))
	if err != nil {
		dialog.NewError(err, appData.mainWindow).Show()
		return
	}
	luri := listable{URI: uri}
	var fdialog *dialog.FileDialog
	if isJSON {
		fdialog = dialog.NewFileOpen(doLoadJSON, appData.mainWindow)
		fdialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	} else {
		fdialog = dialog.NewFileOpen(doLoad, appData.mainWindow)
		fdialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".help"}))
	}
	fdialog.SetLocation(luri)

	fdialog.Show()
}

// doLoad reads the usage text from the supplied path and gets a
// map of boa.CmdLineItems.  This is used for editing a formerly
// used set of commands/flags or resuming an unfinished set.
func doLoad(uwc fyne.URIReadCloser, e error) {
	if uwc == nil {
		return
	}
	defer uwc.Close()

	if e != nil {
		dlg := dialog.NewError(e, appData.mainWindow)
		dlg.Show()
	}

	stat, err := os.Stat(uwc.URI().Path())
	if err != nil {
		dlg := dialog.NewError(err, appData.mainWindow)
		dlg.Show()
		return
	}
	sz := stat.Size()

	b := make([]byte, sz+32)
	_, err = uwc.Read(b)
	if err != nil {
		dlg := dialog.NewError(e, appData.mainWindow)
		dlg.Show()
		return
	}
	b = bytes.Trim(b, "\x00")
	var imap map[string]boa.CmdLineItem
	var errstr, app string

	// boa will read the script and give us a map of allowable
	// command line commands, parameters, flags
	imap, errstr, app = boa.CollectItems(string(b))
	if len(errstr) > 0 {
		e := errors.New(errstr)
		dlg := dialog.NewError(e, appData.mainWindow)
		dlg.Show()
		return
	}
	resetAll()
	appData.appName.SetText(app)

	// sort the items by id so the names can be stored in the same
	// order they were read in
	var items []boa.CmdLineItem
	var ids []int
	for _, i := range imap {
		ids = append(ids, i.Id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		for _, it := range imap {
			if it.Id == id {
				items = append(items, imap[it.Name])
				break
			}
		}
	}

	firstFlag := 0
	firstCmd := 0
	for _, c := range items {
		if c.IsFlag {
			flagInfo.loadFlag(c)
			if firstFlag == 0 {
				firstFlag = c.Id
				flagInfo.populateTab(c.Name)
			}

		} else {
			cmdInfo.loadCommand(c)
			if firstCmd == 0 {
				firstCmd = c.Id
				cmdInfo.populateTab(c.Name)
			}
		}
	}
	refreshPreview()
	appData.loadedPath = uwc.URI().Path()
	appData.dirty = false
}

// doLoadJSON reads the json from the supplied path and gets a
// map of boa.CmdLineItems.  This is used for editing a formerly
// used set of commands/flags or resuming an unfinished set.
func doLoadJSON(uwc fyne.URIReadCloser, e error) {
	if uwc == nil {
		return
	}
	defer uwc.Close()

	if e != nil {
		dlg := dialog.NewError(e, appData.mainWindow)
		dlg.Show()
	}

	stat, err := os.Stat(uwc.URI().Path())
	if err != nil {
		dlg := dialog.NewError(err, appData.mainWindow)
		dlg.Show()
		return
	}
	sz := stat.Size()

	b := make([]byte, sz+32)
	_, err = uwc.Read(b)
	if err != nil {
		dlg := dialog.NewError(e, appData.mainWindow)
		dlg.Show()
		return
	}

	b = bytes.Trim(b, "\x00")
	var imap map[string]boa.CmdLineItem

	// boa will read the json and give us a map of allowable
	// command line commands, parameters, flags
	imap, err = boa.CollectItemsFromJSON(string(b))
	if err != nil {
		dlg := dialog.NewError(e, appData.mainWindow)
		dlg.Show()
		return
	}

	resetAll()
	if appdata, ok := imap[appDataName]; ok {
		appData.appName.SetText(appdata.Alias)
		appData.pkg.SetText(appdata.ShortHelp)
		appData.author.SetText(appdata.LongHelp)
		moreInfo.moreSec.SetText(appdata.Extra)
		delete(imap, appDataName)
	}

	// sort the items by id so the items can be stored in the same
	// order they were read in
	var items []boa.CmdLineItem
	var ids []int
	for _, i := range imap {
		ids = append(ids, i.Id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		for _, it := range imap {
			if it.Id == id {
				items = append(items, imap[it.Name])
				break
			}
		}
	}

	firstFlag := -1
	firstCmd := -1
	for _, c := range items {
		if c.IsFlag {
			flagInfo.loadFlag(c)
			if firstFlag == -1 {
				firstFlag = c.Id
				flagInfo.populateTab(c.Name)
			}
		} else {
			cmdInfo.loadCommand(c)
			if firstCmd == -1 {
				firstCmd = c.Id
				cmdInfo.populateTab(c.Name)
			}
		}
	}
	refreshPreview()
	appData.loadedPath = uwc.URI().Path()
	appData.dirty = false
}

// func reloadFile() {
// 	if appData.loadedPath == "" {
// 		info := dialog.NewInformation("Reload file", "No file is currently loaded.", appData.mainWindow)
// 		info.Show()
// 		return
// 	}
// 	uri := storage.NewFileURI(appData.loadedPath)
// 	reader, err := storage.Reader(uri)
// 	doLoad(reader, err)
// }
