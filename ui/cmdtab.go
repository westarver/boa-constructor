package ui

import (
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	dlg "github.com/sqweek/dialog"
	"github.com/westarver/boa"
	data "github.com/westarver/boa-constructor/appdata"
	"github.com/westarver/boa-constructor/clargs"
	con "github.com/westarver/boa-constructor/constants"
	"github.com/westarver/fynewidgets"
)

var workingMap map[string]fyne.Widget = make(map[string]fyne.Widget)

type toolbarAction struct {
	widget.ToolbarAction
	btn *hoverButton
}

func (t *toolbarAction) ToolbarObject() fyne.CanvasObject {
	return t.btn
}

func newToolbarAction(label string, res fyne.Resource, fn func()) *toolbarAction {
	t := &toolbarAction{}
	t.SetIcon(res)
	t.OnActivated = fn
	t.btn = newHoverButton(label, t.Icon, t.OnActivated)
	return t
}

// commandTab will create the widgets to enter commands
// and the help strings associated with each command.
func CommandTab() *fyne.Container {
	//form1-----------------
	label1 := widget.NewLabelWithStyle(con.Label1Text, fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	text1 := widget.NewEntry()
	text1.OnChanged = func(_ string) { data.GlobalData.Dirty = true }
	label2 := widget.NewLabelWithStyle(con.Label2Text, fyne.TextAlignTrailing, fyne.TextStyle{Bold: false})
	text2 := widget.NewEntry()
	text2.OnChanged = func(_ string) { data.GlobalData.Dirty = true }
	label27 := widget.NewLabelWithStyle(con.Label27Text, fyne.TextAlignTrailing, fyne.TextStyle{Bold: false})
	text27 := widget.NewEntry()
	text27.OnChanged = func(_ string) { data.GlobalData.Dirty = true }

	label3 := widget.NewLabelWithStyle(con.Label3Text, fyne.TextAlignTrailing, fyne.TextStyle{Bold: false})
	text3 := widget.NewEntry()
	text3.OnChanged = func(_ string) { data.GlobalData.Dirty = true }
	label4 := widget.NewLabel(con.Label4Text)
	text4 := widget.NewMultiLineEntry()
	text4.SetMinRowsVisible(2)
	text4.OnChanged = func(_ string) { data.GlobalData.Dirty = true }
	form1 := container.NewGridWithColumns(8, label1, text1, label2, text2, label27, text27, layout.NewSpacer(), layout.NewSpacer())
	form15 := container.New(layout.NewFormLayout(), label3, text3, label4, text4)

	//form2----------------
	checkgrp := widget.NewCheckGroup([]string{con.CheckText1, con.CheckText2, con.CheckText3}, func(options []string) {})
	checkgrp.OnChanged = func(s []string) { data.GlobalData.Dirty = true }
	checkgrpbox := container.New(layout.NewVBoxLayout(), checkgrp, layout.NewSpacer())
	numEntry := fynewidgets.NewNumericalEntry()
	numEntry.OnChanged = func(s string) { data.GlobalData.Dirty = true }
	neLabel := widget.NewLabel(con.NELabelText)
	numEntryBox := container.NewHBox(neLabel, numEntry)
	var select2 *widget.Select
	select2Label := widget.NewLabelWithStyle(con.Select2LabelText, fyne.TextAlignTrailing, fyne.TextStyle{Bold: false})
	checkOpt := widget.NewCheck(con.OptCheckText, func(b bool) { data.GlobalData.Dirty = true })
	select1 := widget.NewSelect([]string{con.Select1Text1, con.Select1Text3, con.Select1Text4}, func(val string) {
		data.GlobalData.Dirty = true

		select2.Enable()
		checkOpt.Enable()
		checkOpt.Checked = false

		if val == con.Select1Text3 {
			numEntry.Enable()
			numEntry.FocusGained()
			return
		}

		if val == con.Select1Text1 {
			select2.Selected = ""
			select2.Disable()
			select2.Refresh()
			checkOpt.Disable()
			checkOpt.Checked = false
			checkOpt.Refresh()
		}

		numEntry.Disable()
		numEntry.Text = ""
	})

	select1box := container.NewVBox(checkOpt, numEntryBox, select1, layout.NewSpacer())
	numEntry.Disable()
	select2 = widget.NewSelect([]string{con.Select2Str, con.Select2Int, con.Select2Float, con.Select2Time, con.Select2Duration, con.Select2Date, con.Select2Path, con.Select2URL, con.Select2Email, con.Select2IP, con.Select2Phone}, func(v string) { data.GlobalData.Dirty = true })
	select2box := container.NewVBox(select2, layout.NewSpacer())
	grid1 := container.New(layout.NewHBoxLayout(), checkgrpbox, layout.NewSpacer(), select1box, layout.NewSpacer(), select2Label, select2box, layout.NewSpacer())
	form2 := container.NewVBox(grid1)

	//form3------------------
	label5 := widget.NewLabel(con.Label5Text)
	text5 := widget.NewEntry()
	text5.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	label6 := widget.NewLabel(con.Label6Text)
	text6 := widget.NewEntry()
	text6.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	form3 := container.New(layout.NewFormLayout(), label5, text5, label6, text6)

	// form4-----------------
	label7 := widget.NewLabel(con.Label7Text)
	text7 := widget.NewMultiLineEntry()
	text7.SetMinRowsVisible(6)
	text7.OnChanged = func(s string) { data.GlobalData.Dirty = true }

	form4 := container.New(layout.NewFormLayout(), label7, text7)

	workingMap[con.Label1Text] = text1
	workingMap[con.Label2Text] = text2
	workingMap[label27.Text] = text27
	workingMap[con.Label3Text] = text3
	workingMap[con.Label4Text] = text4
	workingMap[con.Info1Text] = checkgrp
	workingMap[con.OptCheckText] = checkOpt
	workingMap[con.NELabelText] = numEntry
	workingMap[con.Info2Text] = select1
	workingMap[con.Info3Text] = select2
	workingMap[con.Label5Text] = text5
	workingMap[con.Label6Text] = text6
	workingMap[con.Label7Text] = text7

	toolbar := widget.NewToolbar(
		newToolbarAction("Clear Tab", ResourceClear40dpSvg, ClearTab),

		newToolbarAction("Store Command/Flag Ctrl+T", ResourceAdd40dpSvg,
			func() {
				if workingMap[con.Label1Text].(*widget.Entry).Text == "" {
					return
				}
				SaveArg()
			}),

		newToolbarAction("Delete", ResourceRemove40dpSvg,
			func() {
				txt := (workingMap[con.Label1Text]).(*widget.Entry).Text
				if txt == "" {
					return
				}
				data.GlobalData.MainList.Delete(txt)
				if cur, ok := data.GlobalData.MainList.Current(); ok {
					PopulateTab(cur.Name())
				}
			}),
		newToolbarAction("Previous", ResourceArrowback40dpSvg,
			func() {
				if prev, ok := data.GlobalData.MainList.Previous(); ok {
					PopulateTab(prev.Name())
				}
			}),
		newToolbarAction("Next", ResourceArrowforward40dpSvg,
			func() {
				if nxt, ok := data.GlobalData.MainList.Next(); ok {
					PopulateTab(nxt.Name())
				}
			}),
		newToolbarAction("Undo Delete", ResourceUndo40dpSvg,
			func() {
				e := widget.NewEntry()
				fi := widget.NewFormItem("Command/Flag", e)
				dlg := dialog.NewForm("Undo Delete Command", "Ok", "Cancel", []*widget.FormItem{fi}, func(b bool) {
					if b {
						can := data.GlobalData.MainList.Get(e.Text)
						if can != nil {
							if data.GlobalData.MainList.UnDelete(e.Text) {
								PopulateTab(e.Text)
							}
						}
					}
				}, data.GlobalData.MainWindow)
				dlg.Show()
			}),
		widget.NewToolbarSeparator(),
		newToolbarAction("File Open", ResourceFileopen40dpSvg, doFileOpen),
		newToolbarAction("File Save  Ctrl+S", ResourceSave40dpSvg, doFileSave),
		newToolbarAction("Save As", ResourceSaveas40dpSvg, doFileSaveAs),
		widget.NewToolbarSeparator(),
		newToolbarAction("Purge Command", ResourceRedremovecircle40dpSvg, doCommandPurge))
	return container.NewVBox(container.NewHBox(toolbar), form1, form15, form2, form3, form4, layout.NewSpacer())
}

func SaveArg() {
	var chk []string
	chk = append(chk, workingMap[con.Info1Text].(*widget.CheckGroup).Selected...)
	var def, excl, req, flg bool
	for _, ch := range chk {
		switch ch {
		case con.CheckText1:
			def = true
		case con.CheckText2:
			excl = true
		case con.CheckText3:
			req = true
		}
	}

	if strings.HasPrefix(workingMap[con.Label1Text].(*widget.Entry).Text, "--") {
		flg = true
	}
	pop := workingMap[con.OptCheckText].(*widget.Check).Checked
	var ct int
	switch workingMap[con.Info2Text].(*widget.Select).Selected {
	case con.Select1Text1:
		ct = 0
	case con.Select1Text3:
		ct, _ = strconv.Atoi(workingMap[con.NELabelText].(*fynewidgets.NumericalEntry).Text)
		if ct < 0 && ct > -100 {
			pop = true
		}
	case con.Select1Text4:
		if pop {
			ct = boa.ZeroOrMore
		} else {
			ct = boa.OneOrMore
		}
	}

	var pt boa.ParameterType
	switch workingMap[con.Info3Text].(*widget.Select).Selected {
	case con.Select2Bool:
		pt = boa.TypeBool
	case con.Select2Str:
		pt = boa.TypeString
		if ct < -1 || ct > 1 {
			pt = boa.TypeStringSlice
		}
	case con.Select2Int:
		pt = boa.TypeInt
		if ct < -1 || ct > 1 {
			pt = boa.TypeIntSlice
		}
	case con.Select2Float:
		pt = boa.TypeFloat
		if ct < -1 || ct > 1 {
			pt = boa.TypeFloatSlice
		}
	case con.Select2Time:
		pt = boa.TypeTime
		if ct < -1 || ct > 1 {
			pt = boa.TypeTimeSlice
		}
	case con.Select2Duration:
		pt = boa.TypeTimeDuration
		if ct < -1 || ct > 1 {
			pt = boa.TypeTimeDurationSlice
		}
	case con.Select2Date:
		pt = boa.TypeDate
		if ct < -1 || ct > 1 {
			pt = boa.TypeDateSlice
		}
	case con.Select2Path:
		pt = boa.TypePath
		if ct < -1 || ct > 1 {
			pt = boa.TypePathSlice
		}
	case con.Select2URL:
		pt = boa.TypeURL
		if ct < -1 || ct > 1 {
			pt = boa.TypeURLSlice
		}
	case con.Select2IP:
		pt = boa.TypeIPv4
		if ct < -1 || ct > 1 {
			pt = boa.TypeIPv4Slice
		}
	case con.Select2Email:
		pt = boa.TypeEmail
		if ct < -1 || ct > 1 {
			pt = boa.TypeEmailSlice
		}
	case con.Select2Phone:
		pt = boa.TypePhone
		if ct < -1 || ct > 1 {
			pt = boa.TypePhoneSlice
		}
	}

	var chnames []string
	if workingMap[con.Label5Text].(*widget.Entry).Text == "" {
		chnames = nil
	} else {
		chnames = strings.Split(workingMap[con.Label5Text].(*widget.Entry).Text, " ")
	}

	cmd := boa.CmdLineItem{
		Name:         workingMap[con.Label1Text].(*widget.Entry).Text,
		Alias:        workingMap[con.Label2Text].(*widget.Entry).Text,
		DefaultValue: workingMap[con.Label27Text].(*widget.Entry).Text,
		ShortHelp:    workingMap[con.Label3Text].(*widget.Entry).Text,
		LongHelp:     workingMap[con.Label4Text].(*widget.Entry).Text,
		ChNames:      chnames,
		ParName:      workingMap[con.Label6Text].(*widget.Entry).Text,
		RunCode:      workingMap[con.Label7Text].(*widget.Entry).Text,
		IsDeleted:    false,
		IsRequired:   req,
		IsDefault:    def,
		IsExclusive:  excl,
		IsFlag:       flg,
		IsParamOpt:   pop,
		ParamType:    pt,
		ParamCount:   ct,
		Value:        nil,
	}

	arg := clargs.NewCommandLineArg(&cmd)
	err := data.GlobalData.MainList.Update(arg)
	if err != nil {
		log.Printf("unable to add or update command/flag: %v", err)
		return
	}
}

func ClearTab() {
	workingMap[con.Label1Text].(*widget.Entry).SetText("")
	workingMap[con.Label2Text].(*widget.Entry).SetText("")
	workingMap[con.Label27Text].(*widget.Entry).SetText("")
	workingMap[con.Label3Text].(*widget.Entry).SetText("")
	workingMap[con.Label4Text].(*widget.Entry).SetText("")
	workingMap[con.Label5Text].(*widget.Entry).SetText("")
	workingMap[con.Label6Text].(*widget.Entry).SetText("")
	workingMap[con.Label7Text].(*widget.Entry).SetText("")
	workingMap[con.Info1Text].(*widget.CheckGroup).SetSelected([]string{})
	workingMap[con.OptCheckText].(*widget.Check).SetChecked(false)
	workingMap[con.NELabelText].(*fynewidgets.NumericalEntry).SetText("")
	workingMap[con.Info2Text].(*widget.Select).SetSelected("")
	workingMap[con.Info3Text].(*widget.Select).SetSelected("")
}

func PopulateTab(name string) {
	item := data.GlobalData.MainList.Get(name)
	if item == nil {
		log.Println(name, "cannot be found.")
		return
	}
	if item.Deleted() {
		log.Println(name, "has been deleted. Undelete to show it")
	}
	ClearTab()

	(workingMap[con.Label1Text]).(*widget.Entry).SetText(item.Name())
	(workingMap[con.Label2Text]).(*widget.Entry).SetText(item.Alias())
	(workingMap[con.Label27Text]).(*widget.Entry).SetText(item.DefaultValue())
	(workingMap[con.Label3Text]).(*widget.Entry).SetText(item.ShortHelp())
	(workingMap[con.Label4Text]).(*widget.Entry).SetText(item.LongHelp())
	(workingMap[con.Label5Text]).(*widget.Entry).SetText(strings.Join(item.Children(), " "))
	(workingMap[con.Label6Text]).(*widget.Entry).SetText(item.Parent())
	(workingMap[con.Label7Text]).(*widget.Entry).SetText(item.RunCode())

	var def, excl, req string
	if item.IsDefault() {
		def = con.CheckText1
	}
	if item.IsExclusive() {
		excl = con.CheckText2
	}
	if item.IsRequired() {
		req = con.CheckText3
	}
	(workingMap[con.Info1Text]).(*widget.CheckGroup).SetSelected([]string{def, excl, req})
	(workingMap[con.OptCheckText]).(*widget.Check).SetChecked(item.IsParamOpt())
	if item.ParamCount() < -1 || item.ParamCount() > 1 {
		(workingMap[con.Info2Text]).(*widget.Select).SetSelected(con.Select1Text4)
	} else if item.ParamCount() == 0 {
		(workingMap[con.Info2Text]).(*widget.Select).SetSelected(con.Select1Text1)
	} else {
		(workingMap[con.Info2Text]).(*widget.Select).SetSelected(con.Select1Text3)
	}
	(workingMap[con.NELabelText]).(*fynewidgets.NumericalEntry).SetText(strconv.FormatInt(int64(item.ParamCount()), 10))
	(workingMap[con.Info3Text]).(*widget.Select).SetSelected(boa.TypeToString(item.ParamType()))
	data.GlobalData.Dirty = false
}

func showFileOpen() (string, error) { // gave up on fyne dialogs. they suck.
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	file, err := dlg.File().Title("Open").Filter("JSON", "json").SetStartDir(path).Load()
	return file, err
}

func showFileSave() (string, error) { // gave up on fyne dialogs. they suck.
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	file, err := dlg.File().Title("Save").Filter("JSON", "json").SetStartDir(path).Save()
	return file, err
}
