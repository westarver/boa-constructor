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
		`boa-constructor is the implementation of the boa-constructor
		GUI app that lets the user lay out the cli of their command line app
		 by filling in text fields, checking boxes and selecting from drop
		 down  selectors.  The names of allowable commands and flags along
		 with the number, type and optional/required status of the parameters
		 to  said commands and flags can be defined.  Where possible validation
		 is done to catch command line input errors by users of the app being
		 defined. During the design stages your work can be saved and recalled
		 for editing. The preferred format is now JSON, although the original
		 format is still viable and not going away any time soon.  Saving and
		 using the original format may still be the best choice if you intend
		 to edit it by hand using a text editor.  The input script format is
		 based on the format used by docopt where the user creates a usage/help
		 text that is parsed and a map created of the command line args actually
		 received.  Docopt requires a more rigidly defined and formatted text.
		 Now that the GUI is available the hand written input script has been
		 largely superseded by the JSON format generated by the GUI.
		 The JSON data generated by boa-constructor can be passed to the Boa
		 package function boa.FromJSON(jsonData string, args []string) from
		 your app and your app will receive a data structure containing the
		 command line args with parameters etc. The same map will be obtained
		 by passing to boa.FromHelp(help string, args []string), a proper inpu
		 script aka usage string aka help string.  Go code to start your app,
		 get and evaluate the commands received, and implement the help command
		 can be generated from the GUI. The current implementation creates three
		 source files, main.go, runner.go and help.go.`
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
