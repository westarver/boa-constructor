package io

import (
	"bytes"
	"errors"
	"log"
	"os"
	"sort"

	"github.com/westarver/boa"
	data "github.com/westarver/boa-constructor/appdata"
	"github.com/westarver/boa-constructor/clargs"
)

func DoSaveJson(path string) error {
	j, err := Args2json(data.GlobalData.AppName.Text,
		data.GlobalData.ImportPath.Text,
		data.GlobalData.HelpPkg.Text,
		data.GlobalData.RunPkg.Text,
		data.GlobalData.Author.Text,
		data.GlobalData.More.Text,
	)
	if err != nil {
		return err
	}

	w, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(j)
	return err
}

// DoLoadJSON reads the json from the supplied path and gets a
// map of boa.CmdLineItems, then storing the data.
func DoLoadJSON(path string) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()

	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	sz := stat.Size()

	b := make([]byte, sz+32)
	_, err = r.Read(b)
	if err != nil {
		return err
	}

	b = bytes.Trim(b, "\x00")

	var imap map[string]boa.CmdLineItem

	// boa will read the json and give us a map of allowable
	// command line commands, parameters, flags
	imap, err = boa.CollectItemsFromJSON(b)
	if err != nil {
		return err
	}
	if imap == nil {
		log.Println("boa returned nil")
		return errors.New("boa unable to create items from json data")
	}

	// read 1st object of json data and if it is the
	// appdata object treat it specially then delete it
	if appdata, ok := imap[boa.AppDataName()]; ok {
		data.GlobalData.AppName.SetText(appdata.Alias)
		data.GlobalData.HelpPkg.SetText(appdata.ShortHelp)
		data.GlobalData.RunPkg.SetText(appdata.Value.(string))
		data.GlobalData.Author.SetText(appdata.LongHelp)
		data.GlobalData.More.SetText(appdata.RunCode)
		data.GlobalData.ImportPath.SetText(appdata.DefaultValue)
		delete(imap, boa.AppDataName())
	}

	// sort the items by id so the items can be stored in the same
	// order they were read in
	var items []boa.CmdLineItem
	var ids []int
	for _, i := range imap {
		ids = append(ids, i.Id)
	}

	sort.Ints(ids)

	for _, id := range ids {
		for k, it := range imap {
			if it.Id == id {
				items = append(items, imap[k])
				break
			}
		}
	}

	data.GlobalData.MainList.ResetAll()

	for _, c := range items {
		cmd := boa.CmdLineItem{
			Name:         c.Name,
			Alias:        c.Alias,
			ShortHelp:    c.ShortHelp,
			LongHelp:     c.LongHelp,
			Errors:       c.Errors,
			RunCode:      c.RunCode,
			IsRequired:   c.IsRequired,
			IsDefault:    c.IsDefault,
			IsExclusive:  c.IsExclusive,
			IsFlag:       c.IsFlag,
			ParamType:    c.ParamType,
			ParamCount:   c.ParamCount,
			Id:           c.Id,
			Value:        nil,
			DefaultValue: c.DefaultValue,
			IsDeleted:    c.IsDeleted,
			ParName:      c.ParName,
			ChNames:      c.ChNames,
		}

		arg := clargs.NewCommandLineArg(&cmd)
		err := data.GlobalData.MainList.Update(arg)
		if err != nil {
			return err
		}
	}

	data.GlobalData.LoadedPath = path
	data.GlobalData.Dirty = false
	return nil
}
