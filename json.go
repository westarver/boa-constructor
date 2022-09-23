package boaconstructor

import (
	"encoding/json"
	"log"

	"github.com/westarver/boa"
)

type sliceWrap struct {
	Commands []boa.CmdLineItem `json:"commands"`
}

func slice2json() []byte {
	var jslice sliceWrap
	// create app data record and put it first.  not a real command
	jslice.Commands = append(jslice.Commands, boa.CmdLineItem{
		Id:        -1,
		Name:      appDataName,
		Alias:     appData.appName.Text,
		ShortHelp: appData.pkg.Text,
		LongHelp:  appData.author.Text,
		Extra:     moreInfo.moreSec.Text,
		Disabled:  true,
	})

	// keep them sorted
	for i, c := range cmdInfo.commands {
		c.Id = i
		jslice.Commands = append(jslice.Commands, c)
	}
	var l = len(jslice.Commands)

	for i, f := range flagInfo.flags {
		f.Id = i + l
		jslice.Commands = append(jslice.Commands, f)
	}

	jsonBytes, err := json.Marshal(&jslice)
	if err != nil {
		log.Printf("unable to marshal commands/flags to JSON: %v", err)
		return nil
	}
	//fmt.Println("\nbytes", string(jsonBytes))
	return jsonBytes
}
