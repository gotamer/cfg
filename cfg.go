// Copyright (c) 2012 The GoTamer Authors. All rights reserved.
// Use of this source code is governed by the MIT license.
// The license can be found at http://www.robotamer.com

// gotamer/cfg is a json configuration package
//     * You define your config items in your application as a struct.
//     * You can save a json template of your struct
//     * You can save runtime modifications to the config
package cfg

import (
	"encoding/json"
	"io/ioutil"
)

// Load gets your config from the json file,
// and fills your struct with the option
func Load(filename string, o interface{}) error {
	b, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(b, &o)
		return err
	}
	return err
}

// Save will save your struct to the given filename,
// this is a good way to create a json template
func Save(filename string, o interface{}) error {
	j, err := json.MarshalIndent(&o, "", "\t")
	if err == nil {
		err = ioutil.WriteFile(filename, j, 0660)
		return err
	}
	return err
}
