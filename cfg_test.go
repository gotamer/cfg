// Copyright (c) 2012 The GoTamer Authors. All rights reserved.
// Use of this source code is governed by the MIT license.
// The license can be found at http://www.robotamer.com 

package cfg

import (
	"testing"
)

const filename = "/tmp/testcfg.json"
const dbpath = "/tmp/my.db"

type Sqlite struct {
	Path string
}

func TestSave(t *testing.T) {

	sql := new(Sqlite)
	sql.Path = dbpath

	if err := Save(filename, sql); err != nil {
		t.Errorf("Config Save Error: %g", err)
	}
}

func TestLoad(t *testing.T) {

	sql := new(Sqlite)

	if err := Load(filename, sql); err != nil {
		t.Errorf("Config Load Error: %g", err)
	}

	if sql.Path != dbpath {
		t.Errorf("Config Path doesn't match: %s", sql.Path)
	}
}
