// Copyright (c) 2012 The GoTamer Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license.
// The license can be found in the README file.

// gotamer/cfg is a json configuration package
// You define your config items in your applications main package as a struct
// gotamer/cfg creates a json template file for you, according to the struct 
// definition in your main application, if it doesn't already exist.
package cfg

import (
	"encoding/json"
	"io/ioutil"
)

const file = "cfg.json"

func Get(c interface{}) (err error) {

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return put(&c, err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		println("Bad json File: ", err)
	}
	return err
}

func put(o interface{}, perr error) error {
	j, err := json.MarshalIndent(&o, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, j, 0664)
	if err != nil {
		return err
	}
	println("\n\tWe created a new config template file for you.")
	println("\tPlease edit the file: ", file)
	println("\tFile mode is set to 664!\n")
	return perr
}