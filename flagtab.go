package boaconstructor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func flagTab() *fyne.Container {
	//form1-----------------
	label1 := widget.NewLabel(Label1TextFlag)
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

	// form4-----------------
	button1 := widget.NewButton(Button1TextFlag, func() {
		flagInfo.saveFlag()
		clear()
		if name, ok := flagInfo.next(); ok {
			flagInfo.populateTab(name)
		}
	})
	button2 := widget.NewButton(Button2TextFlag, func() {
		clear()
		if name, ok := flagInfo.previous(); ok {
			flagInfo.populateTab(name)
		}
	})

	button3 := widget.NewButton(Button3TextFlag, func() {
		clear()
	})
	button4 := widget.NewButton(Button4Text, func() {
		saveToFile(true)
	})

	form4 := container.NewHBox(button1, button2, button3, button4, layout.NewSpacer())

	//-------------------------
	divider := widget.NewSeparator()

	flagInfo.working[Label1TextFlag] = text1
	flagInfo.working[Label2Text] = text2
	flagInfo.working[Label3Text] = text3
	flagInfo.working[Label4Text] = text4
	flagInfo.working[Info1Text] = checkgrp
	flagInfo.working[OptCheckText] = checkOpt
	flagInfo.working[NELabelText] = numEntry
	flagInfo.working[Info2TextFlag] = select1
	flagInfo.working[Info3TextFlag] = select2
	flagInfo.working[Label5Text] = text5
	flagInfo.working[Label6Text] = text6

	return container.NewVBox(form1, form2, form3, layout.NewSpacer(), divider, form4)
}
