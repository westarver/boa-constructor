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
// in the command tab

type cmdInfoStore struct {
	commands []boa.CmdLineItem
	current  boa.CmdLineItem
	working  map[string]fyne.Widget
	cursor   int
}

func (c *cmdInfoStore) next() (string, bool) {
	if len(c.commands) == 0 {
		return "", false
	}
	c.cursor++
	if c.cursor >= len(c.commands) {
		c.cursor = 0
	}
	result := c.commands[c.cursor].Name
	return result, true
}

func (c *cmdInfoStore) previous() (string, bool) {
	if len(c.commands) == 0 {
		return "", false
	}
	c.cursor--
	if c.cursor < 0 {
		c.cursor = len(c.commands) - 1
	}
	result := c.commands[c.cursor].Name
	return result, true
}

func (c *cmdInfoStore) get(name string) (*boa.CmdLineItem, bool, int) {
	for i, c := range c.commands {
		if c.Name == name {
			return &c, true, i
		}
	}
	return nil, false, -1
}

func (c *cmdInfoStore) reset() {
	c.current = boa.CmdLineItem{}
}

func (c *cmdInfoStore) resetAll() {
	c.reset()
	c.commands = []boa.CmdLineItem{}
	c.clearTab()
}

func (c *cmdInfoStore) clearTab() {
	c.working[Label1Text].(*widget.Entry).SetText("")
	c.working[Label2Text].(*widget.Entry).SetText("")
	c.working[Label3Text].(*widget.Entry).SetText("")
	c.working[Label4Text].(*widget.Entry).SetText("")
	c.working[Label5Text].(*widget.Entry).SetText("")
	c.working[Label6Text].(*widget.Entry).SetText("")
	c.working[Label7Text].(*widget.Entry).SetText("")
	c.working[Info1Text].(*widget.CheckGroup).SetSelected([]string{})
	c.working[OptCheckText].(*widget.Check).SetChecked(false)
	c.working[NELabelText].(*numericalEntry).SetText("")
	c.working[Info2Text].(*widget.Select).SetSelected("")
	c.working[Info3Text].(*widget.Select).SetSelected("")
}

func (c *cmdInfoStore) populateTab(name string) {
	item, ok, _ := cmdInfo.get(name)
	if !ok {
		return
	}
	c.working[Label1Text].(*widget.Entry).SetText(item.Name)
	c.working[Label2Text].(*widget.Entry).SetText(item.Alias)
	c.working[Label3Text].(*widget.Entry).SetText(item.ShortHelp)
	c.working[Label4Text].(*widget.Entry).SetText(item.LongHelp)
	c.working[Label5Text].(*widget.Entry).SetText(strings.Join(item.RequiredAnd, " "))
	c.working[Label6Text].(*widget.Entry).SetText(strings.Join(item.RequiredOr, " "))
	c.working[Label7Text].(*widget.Entry).SetText(item.Extra)

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
	c.working[Info1Text].(*widget.CheckGroup).SetSelected([]string{def, excl, req})
	c.working[OptCheckText].(*widget.Check).SetChecked(item.ParamOpt)
	switch item.ParamCount {
	case -1:
		c.working[Info2Text].(*widget.Select).SetSelected(Select1Text4)
	case 0:
		c.working[Info2Text].(*widget.Select).SetSelected(Select1Text1)
	default:
		c.working[Info2Text].(*widget.Select).SetSelected(Select1Text3)
	}
	c.working[NELabelText].(*numericalEntry).SetText(strconv.FormatInt(int64(item.ParamCount), 10))
	c.working[Info3Text].(*widget.Select).SetSelected(boa.TypeToString(item.ParamType))
}

func (c *cmdInfoStore) loadCommand(item boa.CmdLineItem) {
	_, _, ndx := c.get(item.Name)
	if ndx != -1 {
		c.commands[ndx] = item
		return
	}
	c.commands = append(c.commands, item)
}

func (c *cmdInfoStore) saveCommand() {
	c.storeName(c.working[Label1Text].(*widget.Entry).Text)
	c.storeAlias(c.working[Label2Text].(*widget.Entry).Text)
	c.storeShort(c.working[Label3Text].(*widget.Entry).Text)
	c.storeLong(c.working[Label4Text].(*widget.Entry).Text)
	c.storeFunction(c.working[Label7Text].(*widget.Entry).Text)

	sl := strings.Split(c.working[Label5Text].(*widget.Entry).Text, " ")
	c.storeReqAnd(sl)

	sl = strings.Split(c.working[Label6Text].(*widget.Entry).Text, " ")
	c.storeReqOr(sl)

	c.storeIsFlag(false)

	var chk []string
	chk = append(chk, c.working[Info1Text].(*widget.CheckGroup).Selected...)
	for _, ch := range chk {
		switch ch {
		case CheckText1:
			c.storeDefault(true)
		case CheckText2:
			c.storeExclusive(true)
		case CheckText3:
			c.storeRequired(true)
		}
	}

	ck := c.working[OptCheckText].(*widget.Check).Checked
	c.storeParamOptional(ck)

	var ct int
	switch c.working[Info2Text].(*widget.Select).Selected {
	case Select1Text1:
		ct = 0
	case Select1Text3:
		ct, _ = strconv.Atoi(c.working[NELabelText].(*numericalEntry).Text)
	case Select1Text4:
		ct = -1
	}
	c.storeParamCount(ct)

	var pt boa.ParameterType
	switch c.working[Info3Text].(*widget.Select).Selected {
	case Select2Bool:
		pt = boa.TypeBool
	case Select2Str:
		pt = boa.TypeString
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeStringSlice
		}
	case Select2Int:
		pt = boa.TypeInt
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeIntSlice
		}
	case Select2Float:
		pt = boa.TypeFloat
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeFloatSlice
		}
	case Select2Time:
		pt = boa.TypeTime
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeTimeSlice
		}
	case Select2Duration:
		pt = boa.TypeTimeDuration
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeTimeDurationSlice
		}
	case Select2Date:
		pt = boa.TypeDate
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeDateSlice
		}
	case Select2Path:
		pt = boa.TypePath
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypePathSlice
		}
	case Select2URL:
		pt = boa.TypeURL
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeURLSlice
		}
	case Select2IP:
		pt = boa.TypeIPv4
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeIPv4Slice
		}
	case Select2Email:
		pt = boa.TypeEmail
		if c.current.ParamCount == -1 || c.current.ParamCount > 1 {
			pt = boa.TypeEmailSlice
		}
	}
	c.storeParamType(pt)
	c.loadCommand(c.current)
	refreshPreview()
	c.reset()
}

func (c *cmdInfoStore) storeName(name string) {
	c.current.Name = name
}

func (c *cmdInfoStore) storeAlias(alias string) {
	c.current.Alias = alias
}

func (c *cmdInfoStore) storeShort(short string) {
	c.current.ShortHelp = short
}

func (c *cmdInfoStore) storeLong(long string) {
	c.current.LongHelp = long
}

func (c *cmdInfoStore) storeFunction(fn string) {
	c.current.Extra = fn
}

func (c *cmdInfoStore) storeReqAnd(req []string) {
	c.current.RequiredAnd = req
}

func (c *cmdInfoStore) storeReqOr(req []string) {
	c.current.RequiredOr = req
}

func (c *cmdInfoStore) storeRequired(req bool) {
	c.current.Required = req
}

func (c *cmdInfoStore) storeExclusive(ex bool) {
	c.current.Exclusive = ex
}

func (c *cmdInfoStore) storeDefault(def bool) {
	c.current.IsDefault = def
}

func (c *cmdInfoStore) storeIsFlag(fl bool) {
	c.current.IsFlag = fl
}

func (c *cmdInfoStore) storeParamType(pt boa.ParameterType) {
	c.current.ParamType = pt
}

func (c *cmdInfoStore) storeParamCount(ct int) {
	c.current.ParamCount = ct
}

func (c *cmdInfoStore) storeParamOptional(po bool) {
	c.current.ParamOpt = po
}
