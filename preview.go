package boaconstructor

import (
	"regexp"
	"strings"

	"github.com/bitfield/script"
)

const previewSize = 16

func previewString() string {
	cmdSec := cmdInfo.commandSection()
	flagSec := flagInfo.flagSection()
	longSec := longHelpSection()
	moreSec := moreInfo.moreSection()
	result := cmdSec + flagSec + longSec + moreSec
	if len(result) == 0 {
		result = strings.Repeat("\n", previewSize)
	}

	return result
}

func refreshPreview() {
	appData.preview.SetText(previewString())
}

func clean() string {
	prev := previewString()
	slice, err := script.Echo(prev).Slice()
	if err != nil {
		return prev
	}

	const metaPat = `[#+*%&]`
	meta := regexp.MustCompile(metaPat)
	for i := 0; i < len(slice); i++ {
		if strings.HasPrefix(slice[i], "Long Description") {
			break
		}
		loc := meta.FindAllStringIndex(slice[i], -1)
		if loc != nil {
			lng := loc[len(loc)-1][1]
			pad := strings.Repeat(" ", lng) + ":"
			slice[i] = slice[i][lng:]
			// the regexp pattern does not match the dot used to denote a float parameter
			// get rid of that here
			slice[i] = strings.TrimPrefix(slice[i], ".")
			slice[i] = strings.Replace(slice[i], ":", pad, 1)
		}
	}

	return strings.Join(slice, "\n")
}
