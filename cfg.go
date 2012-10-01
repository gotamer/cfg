// Copyright (c) 2012 The GoTamer Authors. All rights reserved.
// Use of this source code is governed by the MIT license.
// The license can be found at http://www.robotamer.com 

// gotamer/cfg is a json configuration package
// You define your config items in your application as a struct.  
// If the config file doesn't exist, gotamer/cfg will create a 
// json template file for you, according to the struct 
// definition in your application.
// Usage: 
//		type Sqlite struct {
//			Path string
//		}
//		cfg.APPL = "CafeMaker"
//		cfg.Name = "database"
//		cfgsql := new(Sqlite)
//		err := cfg.Get(cfgsql)
//		fmt.Println(sqlcfg.Path)	
package cfg

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Name is optional, without name the file will be $GOPATH/etc/[APPL].json
// With Name the file will be $GOPATH/etc/[APPL]/[Name].json
var (
	GOPATH = os.Getenv("GOPATH")
	APPL   string // What ever the name of your main package is.
	Name   string // Name for the file to be used without extention
)

// Get gets your config and fills your struct with the option from the json file
func Get(o interface{}) (err error) {
	b, err := ioutil.ReadFile(path())
	if err != nil {
		return createTemplate(&o)
	}
	err = json.Unmarshal(b, &o)
	return
}

// Save will save your runtime modified options to the file. 
// This is mainly created for testing but is exported in case you needed.
func Save(o interface{}) error {
	j, err := json.MarshalIndent(&o, "", "\t")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(path(), j, 0660); err != nil {
		return err
	}
	return nil
}

func createTemplate(o interface{}) (err error) {
	p := mkpath()
	if err = Save(o); err == nil {
		println("\n\tWe created a new config template file for you.")
		println("\tFile mode is set to 660!\n")
		println("\tPlease edit the file: \n", p)
	}
	return err
}

var basepath string

func mkpath() (p string) {
	p = path()
	os.MkdirAll(basepath, 0770)
	return
}

func path() (p string) {
	const ps = "/"
	if APPL == "" {
		println("")
		println("Please set your application name first. cfg.APPL = \"CafeMaker\"")
		println("")
		os.Exit(-1)
	} else {
		if Name == "" {
			basepath = GOPATH + ps + "etc"
			p = basepath + ps + APPL + ".json"
		} else {
			basepath = GOPATH + ps + "etc" + ps + APPL
			p = basepath + ps + Name + ".json"
		}
	}
	return
}
