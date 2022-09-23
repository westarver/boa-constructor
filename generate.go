package boaconstructor

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/westarver/boa"
)

const (
	Label1Text      = "Command"
	Label1TextFlag  = "Flag"
	Label2Text      = "Alias"
	Label3Text      = "Short Help"
	Label4Text      = "Long Help"
	Label5Text      = "Required With Other Commands/Flags"
	Label6Text      = "Required Or Other Commands/Flags"
	Label7Text      = "Code to Execute"
	Button1Text     = "Store/Next Command"
	Button1TextFlag = "Store/Next Flag"
	Button2Text     = "Previous Command"
	Button2TextFlag = "Previous Flag"
	Button3Text     = "Clear Command"
	Button3TextFlag = "Clear Flag"
	Button4Text     = "Save"
	InfoLabel1Text  = "Extra Data"
	Info1Text       = "Options "
	Info2Text       = "Command Parameter(s)"
	Info2TextFlag   = "Flag Parameter(s)"
	Info3Text       = "Command Parameter Type"
	Info3TextFlag   = "Flag Parameter Type"
	CheckText1      = "Default"
	CheckText2      = "Exclusive"
	CheckText3      = "Required"
	NELabelText     = "Fixed Number"
	Select1Text1    = "None"
	OptCheckText    = "Optional"
	Select1Text3    = "Fixed Number"
	Select1Text4    = "Variable Number"
	Select2Bool     = "Bool"
	Select2Str      = "String"
	Select2Int      = "Integer"
	Select2Float    = "Float"
	Select2Time     = "Time"
	Select2Duration = "Time Duration"
	Select2Date     = "Date"
	Select2Path     = "File Path"
	Select2URL      = "URL"
	Select2IP       = "IP Address"
	Select2Email    = "Email Address"
)

var cmdInfo = cmdInfoStore{
	commands: []boa.CmdLineItem{},
	current:  boa.CmdLineItem{},
	working:  make(map[string]fyne.Widget),
	cursor:   0,
}

var flagInfo = flagInfoStore{
	flags:   []boa.CmdLineItem{},
	current: boa.CmdLineItem{},
	working: make(map[string]fyne.Widget),
	cursor:  0,
}

var moreInfo = moreInfoStore{}

type numericalEntry struct {
	widget.Entry
}

func newNumericalEntry() *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *numericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
		e.Entry.TypedRune(r)
	}
}

func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

type readOnlyEntry struct {
	widget.Entry
}

func newreadOnlyEntry() *readOnlyEntry {
	entry := &readOnlyEntry{}
	entry.ExtendBaseWidget(entry)
	entry.MultiLine = true
	entry.TextStyle.Monospace = true
	return entry
}

func (e *readOnlyEntry) TypedRune(r rune) {}

func (e *readOnlyEntry) TypedShortcut(shortcut fyne.Shortcut) {}
