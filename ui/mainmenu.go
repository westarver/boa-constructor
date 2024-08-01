package ui

import (
	"errors"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	dlg "github.com/sqweek/dialog"
	data "github.com/westarver/boa-constructor/appdata"
	gen "github.com/westarver/boa-constructor/generate"
	"github.com/westarver/boa-constructor/help"
	"github.com/westarver/boa-constructor/io"
)

const (
	//menu item strings
	// Opens a json definition file
	FileItemOpen = "Open                     Ctrl+O"
	// Saves the current state to a json definition file
	FileItemSave = "Save                       Ctrl+S"
	// Saves the current state to a different json definition file
	FileItemSaveAs = "Save As       Ctrl+Shift+S"
	// Generates source code from the current state
	FileItemGen = "Generate Code   Ctrl+G"
	// Clear all tabs and start fresh with verification if state changed
	EditItemClear = "Clear"
	// Refresh the preview screen
	EditItemPrev = "Refresh Preview               Ctrl+R"
	// Copy Help Preview to the system clipboard
	EditItemCopy = "Copy help to clipboard"
	// Copy the JSON formatted content to the system clipboard
	EditItemCopyJSON = "Copy JSON to clipboard  Ctrl+L"
	// Brief description of the app
	HelpItemAbout = "About"
	// View the more extensive help
	HelpItemHelp = "Help"
)

func init() {
	CtrlT := &desktop.CustomShortcut{KeyName: fyne.KeyT, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlS := &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlShiftS := &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl | fyne.KeyModifierShift}
	CtrlO := &desktop.CustomShortcut{KeyName: fyne.KeyO, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlR := &desktop.CustomShortcut{KeyName: fyne.KeyR, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlQ := &desktop.CustomShortcut{KeyName: fyne.KeyQ, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlG := &desktop.CustomShortcut{KeyName: fyne.KeyG, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlL := &desktop.CustomShortcut{KeyName: fyne.KeyL, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlRight := &desktop.CustomShortcut{KeyName: fyne.KeyRight, Modifier: fyne.KeyModifierShortcutDefault}
	CtrlLeft := &desktop.CustomShortcut{KeyName: fyne.KeyLeft, Modifier: fyne.KeyModifierShortcutDefault}

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlS, func(shortcut fyne.Shortcut) {
		doFileSave()
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlShiftS, func(shortcut fyne.Shortcut) {
		doFileSaveAs()
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlO, func(shortcut fyne.Shortcut) {
		doFileOpen()
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlR, func(shortcut fyne.Shortcut) {
		RefreshPreview()
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlQ, func(shortcut fyne.Shortcut) {
		data.GlobalData.MainWindow.Close()
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlG, func(shortcut fyne.Shortcut) {
		doFileGen()
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlL, func(shortcut fyne.Shortcut) {
		copyToClipBoardJSON()
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlRight, func(shortcut fyne.Shortcut) {
		if next, ok := data.GlobalData.MainList.Next(); ok {
			PopulateTab(next.Name())
		}
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlLeft, func(shortcut fyne.Shortcut) {
		if prev, ok := data.GlobalData.MainList.Previous(); ok {
			PopulateTab(prev.Name())
		}
	})

	data.GlobalData.MainWindow.Canvas().AddShortcut(CtrlT, func(shortcut fyne.Shortcut) {
		SaveArg()
	})
}

func fileItemOpen() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemOpen, doFileOpen)
}

func fileItemSave() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemSave, doFileSave)
}

func fileItemSaveAs() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemSaveAs, doFileSaveAs)
}
func fileItemGen() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemGen, doFileGen)
}

func editItemClear() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemClear, doEditClear)
}

func editItemPrev() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemPrev, doEditPreview)
}

func editItemCopy() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemCopy, doEditCopy)
}

func editItemCopyJSON() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemCopyJSON, doEditCopyJSON)
}

func helpItemAbout() *fyne.MenuItem {
	return fyne.NewMenuItem(HelpItemAbout, doHelpAbout)
}

func helpItemHelp() *fyne.MenuItem {
	return fyne.NewMenuItem(HelpItemHelp, doHelpHelp)
}

func doFileOpen() {
	file, err := showFileOpen()
	if err != nil {
		log.Println(err)
	}

	err = io.DoLoadJSON(file)
	if err != nil {
		log.Println(err)
		return
	}
	names := data.GlobalData.MainList.Names()
	if len(names) == 0 {
		log.Println(errors.New("DoLoadJSON failed to store any elements"))
		return
	}
	PopulateTab(names[0])

	data.GlobalData.LoadedPath = file
	data.GlobalData.Dirty = false
	RefreshPreview()
}

func doFileSave() {
	file := data.GlobalData.LoadedPath
	if file == "" {
		doFileSaveAs()
		return
	}
	err := io.DoSaveJson(file)
	if err != nil {
		log.Println(err)
	}
	data.GlobalData.Dirty = false
}

func doFileSaveAs() {
	file, err := showFileSave()
	if err != nil {
		log.Println(err)
		return
	}
	err = io.DoSaveJson(file)
	if err != nil {
		log.Println(err)
		return
	}
	data.GlobalData.LoadedPath = file
	data.GlobalData.Dirty = false
}

func doCommandPurge() {
	const confirm = "You are choosing to permanently delete the command/flag %s. Do you want to do it?"
	cur, ok := data.GlobalData.MainList.Current()
	if ok {
		confirmPurge := fmt.Sprintf(confirm, cur.Name())
		yes := dlg.Message("%s", confirmPurge).Title("Delete Permanently").YesNo()
		if yes {
			purged := data.GlobalData.MainList.DeleteAndPurge(cur.Name())
			if purged {
				ClearTab()
				data.GlobalData.Dirty = true
			}
		}
	}
}

func doFileGen() {
	gen.ConfirmGen()
}

func doEditClear() {
	verifyClear()
}

func doEditPreview() {
	RefreshPreview()
}

func doEditCopy() {
	copyToClipBoard()
}

func doEditCopyJSON() {
	copyToClipBoardJSON()
}

func doHelpAbout() {
	about :=
		`Boa Constructor allows you to generate all the code 
and data needed to get and respond to the command line
commands, flags and parameters passed to your cli app.`
	dlg.Message(about).Info()
}

func doHelpHelp() {
	help := help.ShowHelp()

	rt := widget.NewRichTextFromMarkdown(help)
	rt.Resize(fyne.Size{Height: 580, Width: 984})
	rt.Wrapping = fyne.TextWrapWord
	rt.ParseMarkdown(help)
	cont := container.NewVScroll(rt)

	dlg := dialog.NewCustom("Help", "ok", cont, data.GlobalData.MainWindow)
	dlg.Resize(fyne.Size{Height: 620, Width: 1024})
	dlg.Show()
}

func SetMainMenu() {
	file := fyne.NewMenu("File", fileItemOpen(), fileItemSave(), fileItemSaveAs(), fileItemGen())
	edit := fyne.NewMenu("Edit", editItemClear(), editItemPrev(), editItemCopy(), editItemCopyJSON())
	help := fyne.NewMenu("Help", helpItemAbout(), helpItemHelp())
	main := fyne.NewMainMenu(file, edit, help)
	data.GlobalData.MainWindow.SetMainMenu(main)
}

func copyToClipBoard() {
	data.GlobalData.MainWindow.Clipboard().SetContent(PreviewString())
}

func copyToClipBoardJSON() {
	j, err := io.Args2json(data.GlobalData.AppName.Text, data.GlobalData.ImportPath.Text, data.GlobalData.HelpPkg.Text, data.GlobalData.HelpPkg.Text, data.GlobalData.Author.Text, data.GlobalData.More.Text)
	if err != nil {
		return
	}
	data.GlobalData.MainWindow.Clipboard().SetContent(string(j))
}

func verifyClear() {
	if data.GlobalData.Dirty {
		ok := dlg.Message("%s", "The latest edits have not been saved. Do you want to clear?").Title("Work not saved").YesNo()
		if ok {
			ResetAll()
		}
	} else {
		ResetAll()
	}
}
