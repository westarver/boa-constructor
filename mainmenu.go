package boaconstructor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

const (
	// Opens a json definition file
	FileItemOpen string = "Open"
	// Saves the current state to a json definition file
	FileItemSave = "Save"
	// Opens a script definition file
	FileItemOpenLegacy = "Open Script"
	// Saves the current state to a script definition file
	FileItemSaveLegacy = "Save Script"
	// Generates source code from the current  state
	FileItemGen = "Generate Code"
	// Clear all tabs and start fresh with verification if state changed
	EditItemClear = "Clear"
	// Refresh the preview screen
	EditItemPrev = "Refresh Preview"
	// Copy the preview screen to the system clipboard
	EditItemCopy = "Copy script to clipboard"
	// Copy the JSON formatted content to the system clipboard
	EditItemCopyJSON = "Copy JSON to clipboard"
	// Brief description of the app
	HelpItemAbout = "About"
	// View the readme file
	HelpItemHelp = "Help"
)

func ConstString(c string) string {
	switch c {
	case FileItemOpen:
		return "FileItemOpen"
	case FileItemSave:
		return "FileItemSave"
	case FileItemOpenLegacy:
		return "FileItemOpenLegacy"
	case FileItemSaveLegacy:
		return "FileItemSaveLegacy"
	case FileItemGen:
		return "FileItemGen"
	case EditItemClear:
		return "EditItemClear"
	case EditItemPrev:
		return "EditItemPrev"
	case EditItemCopy:
		return "EditItemCopy"
	case EditItemCopyJSON:
		return "EditItemCopyJSON"
	case HelpItemAbout:
		return "HelpItemAbout"
	case HelpItemHelp:
		return "HelpItemHelp"
	}
	return ""
}

func AssocString(c string) string {
	switch c {
	case FileItemOpen:
		return "Opens a json definition file"
	case FileItemSave:
		return "Saves the current state to a json definition file"
	case FileItemOpenLegacy:
		return "Opens a script definition file"
	case FileItemSaveLegacy:
		return "Saves the current state to a script definition file"
	case FileItemGen:
		return "Generates source code from the current  state"
	case EditItemClear:
		return "Clear all tabs and start fresh with verification if state changed"
	case EditItemPrev:
		return "Refresh the preview screen"
	case EditItemCopy:
		return "Copy the preview screen to the system clipboard"
	case EditItemCopyJSON:
		return "Copy the JSON formatted content to the system clipboard"
	case HelpItemAbout:
		return "Brief description of the app"
	case HelpItemHelp:
		return "View the readme file"
	}
	return ""
}

func fileItemOpen() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemOpen, func() { doFileOpen(true) })
}

func fileItemSave() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemSave, func() { doFileSave(true) })
}

func fileItemOpenLegacy() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemOpenLegacy, func() { doFileOpen(false) })
}

func fileItemSaveLegacy() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemSaveLegacy, func() { doFileSave(false) })
}

func fileItemGen() *fyne.MenuItem {
	return fyne.NewMenuItem(FileItemGen, func() { doFileGen() })
}

func editItemClear() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemClear, func() { doEditClear() })
}

func editItemPrev() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemPrev, func() { doEditPreview() })
}

func editItemCopy() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemCopy, func() { doEditCopy() })
}

func editItemCopyJSON() *fyne.MenuItem {
	return fyne.NewMenuItem(EditItemCopyJSON, func() { doEditCopyJSON() })
}

func helpItemAbout() *fyne.MenuItem {
	return fyne.NewMenuItem(HelpItemAbout, func() { doHelpAbout() })
}

func helpItemHelp() *fyne.MenuItem {
	return fyne.NewMenuItem(HelpItemHelp, func() { doHelpHelp() })
}

func doFileOpen(b bool) {
	loadFromFile(b) // true for json, false for legacy script
}

func doFileSave(b bool) {
	saveToFile(b) // true for json, false for legacy script
}

func doFileGen() {
	confirmGen()
}

func doEditClear() {
	verifyClear()
}

func doEditPreview() {
	refreshPreview()
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
	dlg := dialog.NewInformation("About Boa Constructor", about, appData.mainWindow)
	dlg.Show()
}

func doHelpHelp() {
	about :=
		`Boa Constructor allows you to generate all the code 
and data needed to get and respond to the command line
commands, flags and parameters passed to your cli app.`
	dlg := dialog.NewInformation("About Boa Constructor", about, appData.mainWindow)
	dlg.Show()
}

func setMainMenu() {
	file := fyne.NewMenu("File", fileItemOpen(), fileItemSave(), fileItemOpenLegacy(), fileItemSaveLegacy(), fileItemGen())
	edit := fyne.NewMenu("Edit", editItemClear(), editItemPrev(), editItemCopy(), editItemCopyJSON())
	help := fyne.NewMenu("Help", helpItemAbout(), helpItemHelp())
	main := fyne.NewMainMenu(file, edit, help)
	appData.mainWindow.SetMainMenu(main)
}

func copyToClipBoard() {
	appData.mainWindow.Clipboard().SetContent(previewString())
}

func copyToClipBoardJSON() {
	appData.mainWindow.Clipboard().SetContent(string(slice2json()))
}

func verifyClear() {
	if appData.dirty {
		dlg := dialog.NewConfirm("Work not saved", "The latest edits have not been saved. Do you want to clear all?", func(b bool) {
			if b {
				resetAll()
			}
		}, appData.mainWindow)
		dlg.Show()
	} else {
		resetAll()
	}
}
