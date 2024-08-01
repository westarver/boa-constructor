package io

import (
	"encoding/json"
	"errors"

	"github.com/westarver/boa"
	data "github.com/westarver/boa-constructor/appdata"
)

type sliceWrap struct {
	Commands []*boa.CmdLineItem `json:"commands"`
}

// Argsjson takes the slice of boa.CmdLineitem structs
// and outputs it to a JSON formattted byte slice
func Args2json(app, importpath, helppkg, runpkg, auth, more string) ([]byte, error) {
	var jslice sliceWrap

	// create app data record and put it first.  not a real command
	appdata := &boa.CmdLineItem{
		Id:           -1,
		Name:         boa.AppDataName(),
		Alias:        app,        // app name
		ShortHelp:    helppkg,    // store help package name
		Value:        runpkg,     // value will store run package
		DefaultValue: importpath, // save import path
		LongHelp:     auth,
		RunCode:      more,
	}

	if data.GlobalData.MainList.Empty() {
		return nil, errors.New("unable to save to JSON.  no records to save")
	}

	jslice.Commands = append(jslice.Commands, appdata)

	names := data.GlobalData.MainList.Names()
	for i, a := range names {
		arg := data.GlobalData.MainList.Get(a)
		if arg != nil {
			arg.Arg.Id = i
			jslice.Commands = append(jslice.Commands, arg.Arg)
			i++
		}
	}

	jsonBytes, err := json.Marshal(&jslice)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}
