package boaconstructor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// commandTab will create the widgets to enter commands
// and the help strings associated with each command.
func commandTab() *fyne.Container {
	//form1-----------------
	label1 := widget.NewLabel(Label1Text)
	label1.TextStyle = fyne.TextStyle{Bold: true}
	text1 := widget.NewEntry()
	text1.OnChanged = func(s string) { appData.dirty = true }
	label2 := widget.NewLabel(Label2Text)
	label2.TextStyle = fyne.TextStyle{Bold: true}
	text2 := widget.NewEntry()
	text2.OnChanged = func(s string) { appData.dirty = true }
	label3 := widget.NewLabel(Label3Text)
	label3.TextStyle = fyne.TextStyle{Bold: true}
	text3 := widget.NewEntry()
	text3.OnChanged = func(s string) { appData.dirty = true }
	label4 := widget.NewLabel(Label4Text)
	label4.TextStyle = fyne.TextStyle{Bold: true}
	text4 := widget.NewMultiLineEntry()
	text4.SetMinRowsVisible(4)
	text4.OnChanged = func(s string) { appData.dirty = true }
	form1 := container.New(layout.NewFormLayout(), label1, text1, label2, text2, label3, text3, label4, text4)

	//form2----------------
	infoLabel1 := widget.NewLabelWithStyle(InfoLabel1Text, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	info1 := widget.NewLabelWithStyle(Info1Text, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	info2 := widget.NewLabelWithStyle(Info2Text, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	info3 := widget.NewLabelWithStyle(Info3Text, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	checkgrp := widget.NewCheckGroup([]string{CheckText1, CheckText2, CheckText3}, func(options []string) {})
	checkgrp.OnChanged = func(s []string) { appData.dirty = true }
	checkgrpbox := container.New(layout.NewVBoxLayout(), info1, checkgrp, layout.NewSpacer())
	numEntry := newNumericalEntry()
	numEntry.OnChanged = func(s string) { appData.dirty = true }
	neLabel := widget.NewLabel(NELabelText)
	numEntryBox := container.NewHBox(neLabel, numEntry)
	var select2 *widget.Select
	checkOpt := widget.NewCheck(OptCheckText, nil)
	checkOpt.OnChanged = func(b bool) { appData.dirty = true }
	select1 := widget.NewSelect([]string{Select1Text1, Select1Text3, Select1Text4}, func(val string) {
		appData.dirty = true
		select2.Enable()
		checkOpt.Enable()
		checkOpt.Checked = false

		if val == Select1Text3 {
			numEntry.Enable()
			numEntry.FocusGained()
			return
		}

		if val == Select1Text1 {
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

	select1box := container.NewVBox(info2, checkOpt, numEntryBox, select1, layout.NewSpacer())
	numEntry.Disable()
	select2 = widget.NewSelect([]string{Select2Bool, Select2Str, Select2Int, Select2Float, Select2Time, Select2Duration, Select2Date, Select2Path, Select2URL, Select2Email, Select2IP}, func(v string) {})
	select2box := container.NewVBox(info3, select2, layout.NewSpacer())
	grid1 := container.New(layout.NewHBoxLayout(), checkgrpbox, layout.NewSpacer(), select1box, layout.NewSpacer(), select2box, layout.NewSpacer())
	form2 := container.NewVBox(infoLabel1, grid1)

	//form3------------------
	label5 := widget.NewLabel(Label5Text)
	label5.TextStyle = fyne.TextStyle{Bold: true}
	text5 := widget.NewEntry()
	text5.OnChanged = func(s string) { appData.dirty = true }

	label6 := widget.NewLabel(Label6Text)
	label6.TextStyle = fyne.TextStyle{Bold: true}
	text6 := widget.NewEntry()
	text6.OnChanged = func(s string) { appData.dirty = true }

	form3 := container.New(layout.NewFormLayout(), label5, text5, label6, text6)

	// form4-----------------
	label7 := widget.NewLabel(Label7Text)
	label7.TextStyle = fyne.TextStyle{Bold: true}
	text7 := widget.NewMultiLineEntry()
	text7.OnChanged = func(s string) { appData.dirty = true }

	form4 := container.New(layout.NewFormLayout(), label7, text7)

	//clear form
	clear := func() {
		text1.Text = ""
		text1.Refresh()
		text2.Text = ""
		text2.Refresh()
		text3.Text = ""
		text3.Refresh()
		text4.Text = ""
		text4.Refresh()
		text5.Text = ""
		text5.Refresh()
		text6.Text = ""
		text6.Refresh()
		text7.Text = ""
		text7.Refresh()
		checkgrp.Selected = []string{}
		checkgrp.Refresh()
		select1.Selected = ""
		select1.Refresh()
		select2.Selected = ""
		select2.Refresh()
		checkOpt.Checked = false
		checkOpt.Refresh()
		numEntry.Disable()
	}

	// form5-----------------
	button1 := widget.NewButton(Button1Text, func() {
		if text1.Text != "" {
			cmdInfo.saveCommand()
		}
		clear()
		if name, ok := cmdInfo.next(); ok {
			cmdInfo.populateTab(name)
		}
	})

	button2 := widget.NewButton(Button2Text, func() {
		if text2.Text != "" {
			cmdInfo.saveCommand()
		}
		clear()
		if name, ok := cmdInfo.previous(); ok {
			cmdInfo.populateTab(name)
		}
	})
	button3 := widget.NewButton(Button3Text, func() {
		clear()
	})
	button4 := widget.NewButton(Button4Text, func() {
		saveToFile(true)
	})
	form5 := container.NewHBox(button1, button2, button3, button4, layout.NewSpacer())

	//-------------------------
	divider := widget.NewSeparator()

	cmdInfo.working[Label1Text] = text1
	cmdInfo.working[Label2Text] = text2
	cmdInfo.working[Label3Text] = text3
	cmdInfo.working[Label4Text] = text4
	cmdInfo.working[Info1Text] = checkgrp
	cmdInfo.working[OptCheckText] = checkOpt
	cmdInfo.working[NELabelText] = numEntry
	cmdInfo.working[Info2Text] = select1
	cmdInfo.working[Info3Text] = select2
	cmdInfo.working[Label5Text] = text5
	cmdInfo.working[Label6Text] = text6
	cmdInfo.working[Label7Text] = text7

	return container.NewVBox(form1, form2, form3, form4, layout.NewSpacer(), divider, form5)
}
