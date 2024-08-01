package ui

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mgutz/str"
	"github.com/westarver/boa"
	data "github.com/westarver/boa-constructor/appdata"
	"github.com/westarver/boa-constructor/clargs"
)

const previewSize = 14

func PreviewString() string {
	if data.GlobalData.MainList.Empty() {
		return strings.Repeat("\n", previewSize)
	}
	result := "\n\tUsage:\n\n"
	header := headerStr()
	result += "\t" + data.GlobalData.AppName.Text + header + "\n\n"
	cmdSec := argSect()
	result += cmdSec
	longSec := longHelpSection()
	if longSec == "\nLong Descriptions:\n\n" {
		longSec = ""
	}
	result += longSec
	moreSec := data.GlobalData.More.Text
	if moreSec != "" {
		result += "\nMore:\n\n"
		result += moreSec
	}

	if len(result) == 0 {
		result = strings.Repeat("\n", previewSize)
	}
	return result + "\n"
}

func RefreshPreview() {
	data.GlobalData.Preview.SetText(PreviewString() + strings.Repeat("\n", previewSize/3))
}

func argSect() string {
	sb := strings.Builder{}

	sb.WriteString("\n\tCommands/Flags\n\n")
	buf := &bytes.Buffer{}
	names := data.GlobalData.MainList.SortedNames()
	for _, a := range names {
		buf.Reset()
		ar := data.GlobalData.MainList.Get(a)
		if ar == nil {
			break
		}
		if ar.Parent() != "" { // skip sub commands because they are captured
			continue // when the parent is formatted
		}
		arg := format(ar, buf, 0)
		sb.WriteString(arg)
	}
	return sb.String()
}

func yesno(b bool) string {
	return str.Iif(b, "Yes", "No")
}

func format(a *clargs.CommandLineArg, buf *bytes.Buffer, level int) string {
	namesh, col := nameAndShort(a.Name(), a.ShortHelp())
	cttype := paramCtAndType(paramCt(a.ParamCount()), typeStr(a.ParamType()), col)
	var sub, issub bool
	var pop, sn, strbuf string

	if a.Children() != nil {
		sub = true
		sn = fmt.Sprint(a.Children())
		sn = strings.Trim(sn, "[]")
		for _, ch := range a.Children() {
			names := data.GlobalData.MainList.Names()
			for _, n := range names {
				if n != ch {
					continue
				}
				a := data.GlobalData.MainList.Get(n)
				if a == nil {
					continue
				}
				buff := bytes.Buffer{}
				strbuf = format(a, &buff, level+1)
			}
		}
	}

	if a.Parent() != "" {
		issub = true
	}
	if a.ParamCount() == 0 {
		pop = "NA"
	} else {
		pop = yesno(a.IsParamOpt())
	}

	fmt.Fprintf(buf, "	Name:        %s\n", namesh)
	var alias = a.Alias()
	if alias == "" {
		alias = "No alias"
	}
	fmt.Fprintf(buf, "	Alias:       %s\n", alias)
	fmt.Fprintf(buf, "	Param count: %s\n", cttype)
	fmt.Fprintf(buf, "	Param opt:   %s\n", pop)
	if a.Deleted() {
		fmt.Fprintf(buf, "	Deleted:     %v\n", "Yes")
	}
	fmt.Fprintf(buf, "	Exclusive:   %s\n", yesno(a.IsExclusive()))
	fmt.Fprintf(buf, "	Default:     %s\n", yesno(a.IsDefault()))
	fmt.Fprintf(buf, "	Required:    %s\n", yesno(a.IsRequired()))

	if issub {
		fmt.Fprintf(buf, "	Sub command of: %s\n", a.Parent())
	}

	if sub {
		fmt.Fprintf(buf, "	Has sub cmds/flags: %s\n", sn)
	}

	indented := indent(buf.String(), level)

	sb := strings.Builder{}
	sb.WriteString(indented)
	if len(strbuf) != 0 {
		sb.WriteString(strings.Repeat("\t", level+2) + "|\n")
		sb.WriteString(strbuf)
	} else {
		sb.WriteString("\n")
	}
	return sb.String()
}

func indent(str string, level int) string {
	block := strings.Split(str, "\n")
	ind := strings.Repeat("\t", level)
	for i, line := range block {
		block[i] = ind + line
	}
	return strings.TrimRight(strings.Join(block, "\n"), "\n\t\x00 ") + "\n"
}

func paramCtAndType(ct, ty string, col int) string {
	lead := len(ct)
	gap := col - lead
	return ct + strings.Repeat(" ", gap) + ty
}

func paramCt(c int) string {
	if c == 0 {
		return "none"
	}
	if c == -100 {
		return "1 or more"
	}
	if c == -99 {
		return "0 or more"
	}
	ct := fmt.Sprintf("%d", c)
	return strings.Trim(ct, "-")
}

func typeStr(t boa.ParameterType) string {
	switch t {
	case boa.TypeBool:
		return "Boolean"
	case boa.TypeString, boa.TypeStringSlice:
		return "String"
	case boa.TypeInt, boa.TypeIntSlice:
		return "Integer"
	case boa.TypeFloat, boa.TypeFloatSlice:
		return "Float"
	case boa.TypeTime, boa.TypeTimeSlice:
		return "Time"
	case boa.TypeTimeDuration, boa.TypeTimeDurationSlice:
		return "Duration"
	case boa.TypeDate, boa.TypeDateSlice:
		return "Date"
	case boa.TypePath, boa.TypePathSlice:
		return "Path"
	case boa.TypeURL, boa.TypeURLSlice:
		return "URL"
	case boa.TypeIPv4, boa.TypeIPv4Slice:
		return "IPv4"
	case boa.TypeEmail, boa.TypeEmailSlice:
		return "Email"
	case boa.TypePhone, boa.TypePhoneSlice:
		return "Phone"
	}
	return ""
}

const fieldSize = 18

func nameAndShort(n, sh string) (string, int) {
	lenn := len(n)
	gap := fieldSize - lenn
	if gap < 4 {
		gap = 4
	}
	return n + strings.Repeat(" ", gap) + sh, lenn + gap
}

func longHelpSection() string {
	var sb strings.Builder
	sb.WriteString("\nLong Descriptions:\n\n")

	type pair struct {
		leader, lh string
	}

	var lines []pair
	var line string

	names := data.GlobalData.MainList.SortedNames()
	for _, a := range names {
		ar := data.GlobalData.MainList.Get(a)
		if ar == nil {
			break
		}
		line = ar.Name()
		lh := strings.Trim(ar.LongHelp(), ": \t\n")
		if lh == "" {
			continue
		}
		lines = append(lines, pair{leader: line, lh: lh})
	}

	maxlen := 0
	for _, l := range lines {
		if len(l.leader) > maxlen {
			maxlen = len(l.leader)
		}
	}

	helppos := maxlen + 4

	for _, l := range lines {
		var line string
		if len(l.lh) == 0 {
			continue
		}

		padlen := helppos - len(l.leader) - 1 // account for colon character
		line = func(ld, lh string) string {
			lhl := strings.Split(lh, "\n")
			if len(lhl) > 0 {
				lhl[0] = ld + ":" + strings.Repeat(" ", padlen) + lhl[0]
				for i := 1; i < len(lhl); i++ {
					lhl[i] = strings.Repeat(" ", helppos) + lhl[i]
				}
			}
			return strings.Join(lhl, "\n")
		}(l.leader, l.lh)
		sb.WriteString(line + "\n\n")
	}
	return sb.String()
}

func headerStr() string {
	sb := strings.Builder{}
	sb.WriteString(" ")
	var req int
	var flg bool
	list := data.GlobalData.MainList
	names := list.Names()
	for _, n := range names {
		a := list.Get(n)
		if a == nil {
			break
		}
		if a.Arg.IsRequired {
			req++
			dots := a.ParamCount() != -1 || a.ParamCount() != 0 || a.ParamCount() != 1
			parm := boa.TypeToString(a.ParamType()) + str.Iif(dots, "...", "")
			var open, close string
			if a.Arg.IsParamOpt {
				open = " ["
				close = "]"
			} else {
				open = " <"
				close = ">"
			}
			sb.WriteString(a.Name() + open + str.Iif(a.ParamCount() != 0, parm, "") + close)
		}
		if a.Arg.IsFlag {
			flg = true
		}
	}
	if req < len(names)-1 {
		sb.WriteString(" [commands] ")
		if flg {
			sb.WriteString("[flags]")
		}
	}
	return sb.String()
}
