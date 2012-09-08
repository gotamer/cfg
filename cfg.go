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

func Get(c interface{}) (ok error) {

	b, ok := ioutil.ReadFile(file)
	if ok != nil {
		return put(&c, ok)
	}
	ok = json.Unmarshal(b, &c)
	if ok != nil {
		println("Bad json File: ", ok)
	}
	return ok
}

func put(o interface{}, e error) error {
	j, ok := json.MarshalIndent(&o, "", "\t")
	if ok != nil {
		return ok
	}

	ok = ioutil.WriteFile(file, j, 0664)
	if ok != nil {
		return ok
	}
	println("\n\tWe created a new config template file for you.")
	println("\tPlease edit the file: ", file)
	println("\tFile mode is set to 664!\n")
	return e
}