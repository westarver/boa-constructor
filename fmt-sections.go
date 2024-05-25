package boaconstructor

import (
	"fmt"
	"strings"

	"github.com/westarver/boa"
)

// the code in this source file will take the data objects
// created from the GUI input and turn it into an input
// script with meta-characters to denote the various options
// associated with each command/flag, such as number of
// parameters, type of parameters, exclusivity of command,
// required/optional status.  A cleaned version of the input
// script, minus the meta-characters, can then be generated
// to serve as the help text for your cli app.

func (c *cmdInfoStore) commandSection() string {
	var sb strings.Builder
	app := "[app-name]"
	if len(appData.appName.Text) > 0 {
		app = appData.appName.Text
	}
	cmdsec := "Usage: " + app + ` [commands] [flags]
	
Commands:
`
	sb.WriteString(cmdsec)

	type pair struct {
		cmd, sh string
	}
	var lines []pair

	for _, c := range c.commands {
		if !c.Disabled {
			line := formatLine(&c)
			sh := strings.Trim(strings.TrimPrefix(c.ShortHelp, c.Name), "\t ")
			lines = append(lines, pair{cmd: line, sh: sh})
		}
	}

	maxlen := 0
	for _, l := range lines {
		if len(l.cmd) > maxlen {
			maxlen = len(l.cmd)
		}
	}

	colonpos := maxlen + 4
	for _, l := range lines {
		padlen := colonpos - len(l.cmd)
		pad := strings.Repeat(" ", padlen)
		line := l.cmd + pad + ":" + l.sh
		sb.WriteString(line + "\n")
	}

	return sb.String()
}

func (f *flagInfoStore) flagSection() string {
	var sb strings.Builder

	sb.WriteString("\nFlags:\n")

	type pair struct {
		flg, sh string
	}
	var lines []pair

	for _, f := range f.flags {
		if !f.Disabled {
			line := formatLine(&f)
			sh := strings.Trim(strings.TrimPrefix(f.ShortHelp, f.Name), "\t ")
			lines = append(lines, pair{flg: line, sh: sh})
		}
	}

	maxlen := 0
	for _, l := range lines {
		if len(l.flg) > maxlen {
			maxlen = len(l.flg)
		}
	}

	colonpos := maxlen + 4

	for _, l := range lines {
		var line string
		padlen := colonpos - len(l.flg)
		pad := strings.Repeat(" ", padlen)
		line = l.flg + pad + ":" + l.sh
		sb.WriteString(line + "\n")
	}

	return sb.String()
}

func longHelpSection() string {
	var sb strings.Builder
	sb.WriteString("\nLong Description:\n")

	type pair struct {
		leader, lh string
	}

	var lines []pair
	var line string
	for _, c := range cmdInfo.commands {
		line = c.Name
		lh := strings.Trim(c.LongHelp, ": \t\n")
		lines = append(lines, pair{leader: line, lh: lh})
	}

	for _, f := range flagInfo.flags {
		line = f.Name
		lh := strings.Trim(f.LongHelp, ": \t\n")
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

func (m *moreInfoStore) moreSection() string {
	if len(m.moreSec.Text) > 0 {
		return "\nMore:\n" + m.moreSec.Text
	}
	return ""
}

func formatLine(c *boa.CmdLineItem) string {
	var line string
	if c.IsDefault {
		line += "+"
	}
	if c.Exclusive {
		line += "*"
	}
	if c.ParamType == boa.TypeInt || c.ParamType == boa.TypeIntSlice {
		line += "#"
	}
	if c.ParamType == boa.TypeFloat || c.ParamType == boa.TypeFloatSlice {
		line += "."
	}
	if c.ParamType == boa.TypeDate || c.ParamType == boa.TypeDateSlice {
		line += "!"
	}
	if c.ParamType == boa.TypeTime || c.ParamType == boa.TypeTimeSlice {
		line += "%"
	}
	if c.ParamType == boa.TypeTimeDuration || c.ParamType == boa.TypeTimeDurationSlice {
		line += "^"
	}
	if c.ParamType == boa.TypePath || c.ParamType == boa.TypePathSlice {
		line += "/"
	}
	if c.ParamType == boa.TypeURL || c.ParamType == boa.TypeURLSlice {
		line += "\\"
	}
	if c.ParamType == boa.TypeEmail || c.ParamType == boa.TypeEmailSlice {
		line += "@"
	}
	if c.ParamType == boa.TypeIPv4 || c.ParamType == boa.TypeIPv4Slice {
		line += "&"
	}

	if !c.Required {
		line += "["
	}
	line += c.Name
	if c.Alias != "" {
		line += " | " + c.Alias
	}
	if !c.Required {
		line += "]"
	}

	var pre, suf string
	if c.ParamOpt {
		pre = " ["
		suf = "]"
	}

	if c.ParamCount == -1 {
		line += " " + pre + "..." + suf
	}
	if c.ParamCount > 0 {
		pc := fmt.Sprintf("%d", c.ParamCount)
		line += " " + pre + pc + suf
	}
	return line
}
