// Copyright 2018 The go-tstereum Authors
// This file is part of go-tstereum.
//
// go-tstereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-tstereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-tstereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/docker/docker/pkg/reexec"
	"github.com/tstchain/go-tstchain/internal/cmdtest"
)

type testTstkey struct {
	*cmdtest.TestCmd
}

// spawns tstkey with the given command line args.
func runTstkey(t *testing.T, args ...string) *testTstkey {
	tt := new(testTstkey)
	tt.TestCmd = cmdtest.NewTestCmd(t, tt)
	tt.Run("tstkey-test", args...)
	return tt
}

func TestMain(m *testing.M) {
	// Run the app if we've been exec'd as "tstkey-test" in runTstkey.
	reexec.Register("tstkey-test", func() {
		if err := app.Run(os.Args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	})
	// check if we have been reexec'd
	if reexec.Init() {
		return
	}
	os.Exit(m.Run())
}
