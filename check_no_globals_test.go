package main

import (
	"testing"
)

func TestCheckNoGlobals(t *testing.T) {
	cases := []struct {
		path         string
		wantMessages []string
	}{
		{
			path:         "testdata/0",
			wantMessages: nil,
		},
		{
			path:         "testdata/1",
			wantMessages: nil,
		},
		{
			path: "testdata/2",
			wantMessages: []string{
				"testdata/2/code.go:3 myVar is a global variable",
			},
		},
		{
			path: "testdata/3",
			wantMessages: []string{
				"testdata/3/code_0.go:8 theVar is a global variable",
				"testdata/3/code_1.go:3 myVar is a global variable",
			},
		},
		{
			path: "testdata/4",
			wantMessages: []string{
				"testdata/4/code.go:3 theVar is a global variable",
			},
		},
		{
			path: "testdata/4/...",
			wantMessages: []string{
				"testdata/4/code.go:3 theVar is a global variable",
				"testdata/4/subpkg/code_1.go:3 myVar is a global variable",
			},
		},
		{
			path: "testdata/5",
			wantMessages: []string{
				"testdata/5/code.go:3 myVar1 is a global variable",
				"testdata/5/code.go:3 myVar2 is a global variable",
			},
		},
		{
			path:         "testdata/6",
			wantMessages: nil,
		},
		{
			path:         ".",
			wantMessages: nil,
		},
		{
			path: "./...",
			wantMessages: []string{
				"testdata/2/code.go:3 myVar is a global variable",
				"testdata/3/code_0.go:8 theVar is a global variable",
				"testdata/3/code_1.go:3 myVar is a global variable",
				"testdata/4/code.go:3 theVar is a global variable",
				"testdata/4/subpkg/code_1.go:3 myVar is a global variable",
				"testdata/5/code.go:3 myVar1 is a global variable",
				"testdata/5/code.go:3 myVar2 is a global variable",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.path, func(t *testing.T) {
			messages, err := checkNoGlobals(c.path)
			if err != nil {
				t.Fatalf("got error %#v", err)
			}
			if !stringSlicesEqual(messages, c.wantMessages) {
				t.Errorf("got %#v, want %#v", messages, c.wantMessages)
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
