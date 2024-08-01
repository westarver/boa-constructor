package generate

import (
	"strings"

	data "github.com/westarver/boa-constructor/appdata"
)

func LongHelpSection() string {
	var sb strings.Builder
	sb.WriteString("\nLong Description:\n\n")

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
