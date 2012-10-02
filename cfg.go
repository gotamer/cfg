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
	"fmt"
	"io/ioutil"
	"os"
)

// Name is optional, without name the file will be 
//     $GOPATH/etc/[APPL].json
// With Name the file will be 
//     $GOPATH/etc/[APPL]/[Name].json
var (
	ROOT uint8
	APPL string // What ever the name of your main package is.
	Name string // Name for the file to be used without extention
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
		fmt.Println("\n\tWe created a new config template file for you.")
		fmt.Println("\tFile mode is set to 660!\n")
		fmt.Println("\tPlease edit the file: \n", p)
	}
	return err
}

var basepath string

// @todo check write permission for basepath first
func mkpath() (p string) {
	p = path()
	os.MkdirAll(basepath, 0770)
	return
}

func path() (path string) {
	const ps = "/"
	if APPL == "" {
		printError(2)
	}

	switch ROOT {
	case 1, 11:
		basepath = os.Getenv("GOPATH") + "/etc/"
	case 2, 12:
		basepath = os.Getenv("HOME") + "/."
	case 3, 13:
		basepath = os.Getenv("HOME") + "/.config/"
	case 4, 14:
		basepath = os.Getenv("HOME") + "/.go/etc/"
	case 5, 15:
		basepath = "/etc/"
	default:
		printError(1)
	}

	if ROOT < 10 {
		path = basepath + APPL + ".json"
	} else {
		path = basepath + APPL + ps + Name + ".json"
	}
	return
}

func printError(errorno uint8) {
	fmt.Println("")
	fmt.Println("gotamer/cfg ERROR")
	fmt.Println("")
	switch errorno {
	case 1:
		fmt.Println("Please set your ROOT configuration folder.")
		fmt.Println("")
		fmt.Println("\tcfg.ROOT = 1 // $GOPATH/etc/[APPL].json")
		fmt.Println("\tcfg.ROOT = 2 // $HOME/.[APPL].json")
		fmt.Println("\tcfg.ROOT = 3 // $HOME/.config/[APPL].json")
		fmt.Println("\tcfg.ROOT = 4 // $HOME/.go/etc/[APPL].json")
		fmt.Println("\tcfg.ROOT = 5 // /etc/[APPL].json")

		fmt.Println("\tcfg.ROOT = 11 // $GOPATH/etc/[APPL]/Name.json")
		fmt.Println("\tcfg.ROOT = 12 // $HOME/.[APPL]/Name.json")
		fmt.Println("\tcfg.ROOT = 13 // $HOME/.config/[APPL]/Name.json")
		fmt.Println("\tcfg.ROOT = 14 // $HOME/.go/etc/[APPL]/Name.json")
		fmt.Println("\tcfg.ROOT = 15 // /etc/[APPL]/Name.json")
	case 2:
		fmt.Println("Please set your application name first. cfg.APPL = \"CafeMaker\"")
	default:
		fmt.Println("Undefined error.")
	}
	fmt.Println("")
	os.Exit(-1)
}
