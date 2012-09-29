// Copyright (c) 2012 The GoTamer Authors. All rights reserved.
// Use of this source code is governed by the MIT license.
// The license can be found at http://www.robotamer.com 

package cfg

import (
	"testing"
)

func TestGet(t *testing.T) {
	type Sqlite struct {
		Path string
	}
	APPL = "cfgtest"
	Name = "database"
	sql := new(Sqlite)
	if err := Get(sql); err != nil {
		t.Errorf("Config Get 1 Error: %g", err)
	}
	sql.Path = "/tmp"
	if err := Save(sql); err != nil {
		t.Errorf("Config Save Error: %g", err)
	}

	sqltwo := new(Sqlite)
	if err := Get(sqltwo); err != nil {
		t.Errorf("Config Get 2 error : %g", err)
	}
	if sqltwo.Path != "/tmp" {
		t.Errorf("Config should be \tmp : %s", sqltwo.Path)
	}
}
