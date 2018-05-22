package main

import (
	"path/filepath"
	"strconv"
	"testing"
)

func TestCheckNoGlobals(t *testing.T) {
	cases := [][]string{
		nil,
		nil,
		{
			"testdata/2/code.go:3 myVar is a global variable",
		},
		{
			"testdata/3/code_0.go:8 theVar is a global variable",
			"testdata/3/code_1.go:3 myVar is a global variable",
		},
		{
			"testdata/4/subpkg/code_1.go:3 myVar is a global variable",
		},
		{
			"testdata/5/code.go:3 myVar1 is a global variable",
			"testdata/5/code.go:3 myVar2 is a global variable",
		},
		nil,
	}

	for i, wantMessages := range cases {
		testdataName := strconv.FormatInt(int64(i), 10)
		t.Run(testdataName, func(t *testing.T) {
			path := filepath.Join("testdata", testdataName)
			messages, err := checkNoGlobals(path)
			if err != nil {
				t.Fatalf("got error %#v", err)
			}
			if !stringSlicesEqual(messages, wantMessages) {
				t.Errorf("got %#v, want %#v", messages, wantMessages)
			}
		})
	}
}

func stringSlicesEqual(s1, s2 []string) bool {
	diff := map[string]int{}
	for _, s := range s1 {
		diff[s]++
	}
	for _, s := range s2 {
		diff[s]--
		if diff[s] == 0 {
			delete(diff, s)
		}
	}
	return len(diff) == 0
}
