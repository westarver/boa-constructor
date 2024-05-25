package boaconstructor

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/westarver/boa"
)

// this file takes care of the task of creating the structs
// and managing the slice of structs that the user creates
// in the flag tab

type flagInfoStore struct {
	flags   []boa.CmdLineItem
	current boa.CmdLineItem
	working map[string]fyne.Widget
	cursor  int
}

func (f *flagInfoStore) next() (string, bool) {
	if len(f.flags) == 0 {
		return "", false
	}
	f.cursor++
	if f.cursor >= len(f.flags) {
		f.cursor = 0
	}
	result := f.flags[f.cursor].Name
	return result, true
}

func (f *flagInfoStore) previous() (string, bool) {
	if len(f.flags) == 0 {
		return "", false
	}
	f.cursor--
	if f.cursor < 0 {
		f.cursor = len(f.flags) - 1
	}
	result := f.flags[f.cursor].Name
	return result, true
}

func (f *flagInfoStore) get(name string) (*boa.CmdLineItem, bool, int) {
	for i, c := range f.flags {
		if c.Name == name {
			return &c, true, i
		}
	}
	return nil, false, -1
}

func (f *flagInfoStore) reset() {
	f.current = boa.CmdLineItem{}
}

func (f *flagInfoStore) resetAll() {
	f.reset()
	f.flags = []boa.CmdLineItem{}
	f.clearTab()
}

func (f *flagInfoStore) clearTab() {
	f.working[Label1TextFlag].(*widget.Entry).SetText("")
	f.working[Label2Text].(*widget.Entry).SetText("")
	f.working[Label3Text].(*widget.Entry).SetText("")
	f.working[Label4Text].(*widget.Entry).SetText("")
	f.working[Label5Text].(*widget.Entry).SetText("")
	f.working[Label6Text].(*widget.Entry).SetText("")
	f.working[Info1Text].(*widget.CheckGroup).SetSelected([]string{})
	f.working[OptCheckText].(*widget.Check).SetChecked(false)
	f.working[NELabelText].(*numericalEntry).SetText("")
	f.working[Info2TextFlag].(*widget.Select).SetSelected("")
	f.working[Info3TextFlag].(*widget.Select).SetSelected("")
}

func (f *flagInfoStore) populateTab(name string) {
	item, ok, _ := flagInfo.get(name)
	if !ok {
		return
	}
	f.working[Label1TextFlag].(*widget.Entry).SetText(item.Name)
	f.working[Label2Text].(*widget.Entry).SetText(item.Alias)
	f.working[Label3Text].(*widget.Entry).SetText(item.ShortHelp)
	f.working[Label4Text].(*widget.Entry).SetText(item.LongHelp)
	f.working[Label5Text].(*widget.Entry).SetText(strings.Join(item.RequiredAnd, " "))
	f.working[Label6Text].(*widget.Entry).SetText(strings.Join(item.RequiredOr, " "))

	var def, excl, req string
	if item.IsDefault {
		def = CheckText1
	}
	if item.Exclusive {
		excl = CheckText2
	}
	if item.Required {
		req = CheckText3
	}
	f.working[Info1Text].(*widget.CheckGroup).SetSelected([]string{def, excl, req})
	f.working[OptCheckText].(*widget.Check).SetChecked(item.ParamOpt)
	switch item.ParamCount {
	case -1:
		f.working[Info2TextFlag].(*widget.Select).SetSelected(Select1Text4)
	case 0:
		f.working[Info2TextFlag].(*widget.Select).SetSelected(Select1Text1)
	default:
		f.working[Info2TextFlag].(*widget.Select).SetSelected(Select1Text3)
	}
	f.working[NELabelText].(*numericalEntry).SetText(strconv.FormatInt(int64(item.ParamCount), 10))
	f.working[Info3TextFlag].(*widget.Select).SetSelected(boa.TypeToString(item.ParamType))
}

func (f *flagInfoStore) loadFlag(item boa.CmdLineItem) {
	_, _, ndx := flagInfo.get(item.Name)
	if ndx != -1 {
		f.flags[ndx] = item
		return
	}
	f.flags = append(f.flags, item)
}

func (f *flagInfoStore) saveFlag() {
	f.storeName(f.working[Label1TextFlag].(*widget.Entry).Text)
	f.storeAlias(f.working[Label2Text].(*widget.Entry).Text)
	f.storeShort(f.working[Label3Text].(*widget.Entry).Text)
	f.storeLong(f.working[Label4Text].(*widget.Entry).Text)

	sl := strings.Split(f.working[Label5Text].(*widget.Entry).Text, " ")
	f.storeReqAnd(sl)

	sl = strings.Split(f.working[Label6Text].(*widget.Entry).Text, " ")
	f.storeReqOr(sl)

	f.storeIsFlag(true)

	var chk []string
	chk = append(chk, f.working[Info1Text].(*widget.CheckGroup).Selected...)
	for _, ch := range chk {
		switch ch {
		case CheckText1:
			f.storeDefault(true)
		case CheckText2:
			f.storeExclusive(true)
		case CheckText3:
			f.storeRequired(true)
		}
	}

	ck := f.working[OptCheckText].(*widget.Check).Checked
	f.storeParamOptional(ck)

	var ct int
	switch f.working[Info2TextFlag].(*widget.Select).Selected {
	case Select1Text1:
		ct = 0
	case Select1Text3:
		ct, _ = strconv.Atoi(f.working[NELabelText].(*numericalEntry).Text)
	case Select1Text4:
		ct = -1
	}
	f.storeParamCount(ct)

	var pt boa.ParameterType
	switch f.working[Info3TextFlag].(*widget.Select).Selected {
	case Select2Bool:
		pt = boa.TypeBool
	case Select2Str:
		pt = boa.TypeString
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeStringSlice
		}
	case Select2Int:
		pt = boa.TypeInt
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeIntSlice
		}
	case Select2Float:
		pt = boa.TypeFloat
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeFloatSlice
		}
	case Select2Time:
		pt = boa.TypeTime
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeTimeSlice
		}
	case Select2Duration:
		pt = boa.TypeTimeDuration
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeTimeDurationSlice
		}
	case Select2Date:
		pt = boa.TypeDate
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeDateSlice
		}
	case Select2Path:
		pt = boa.TypePath
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypePathSlice
		}
	case Select2URL:
		pt = boa.TypeURL
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeURLSlice
		}
	case Select2IP:
		pt = boa.TypeIPv4
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeIPv4Slice
		}
	case Select2Email:
		pt = boa.TypeEmail
		if f.current.ParamCount == -1 || f.current.ParamCount > 1 {
			pt = boa.TypeEmailSlice
		}
	}
	f.storeParamType(pt)
	f.loadFlag(f.current)
	refreshPreview()
	f.reset()
}

func (f *flagInfoStore) storeName(name string) {
	f.current.Name = name
}

func (f *flagInfoStore) storeAlias(alias string) {
	f.current.Alias = alias
}

func (f *flagInfoStore) storeShort(short string) {
	f.current.ShortHelp = short
}

func (f *flagInfoStore) storeLong(long string) {
	f.current.LongHelp = long
}

func (f *flagInfoStore) storeReqAnd(req []string) {
	f.current.RequiredAnd = req
}

func (f *flagInfoStore) storeReqOr(req []string) {
	f.current.RequiredOr = req
}

func (f *flagInfoStore) storeRequired(req bool) {
	f.current.Required = req
}

func (f *flagInfoStore) storeExclusive(ex bool) {
	f.current.Exclusive = ex
}

func (f *flagInfoStore) storeDefault(def bool) {
	f.current.IsDefault = def
}

func (f *flagInfoStore) storeIsFlag(fl bool) {
	f.current.IsFlag = fl
}

func (f *flagInfoStore) storeParamType(pt boa.ParameterType) {
	f.current.ParamType = pt
}

func (f *flagInfoStore) storeParamCount(ct int) {
	f.current.ParamCount = ct
}

func (f *flagInfoStore) storeParamOptional(po bool) {
	f.current.ParamOpt = po
}
