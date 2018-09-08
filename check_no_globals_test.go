package main

import (
	"fmt"
	"testing"
)

func TestCheckNoGlobals(t *testing.T) {
	cases := []struct {
		path         string
		includeTests bool
		wantMessages []string
	}{
		{
			path:         "testdata/0",
			wantMessages: nil,
		},
		{
			path:         "testdata/0",
			includeTests: true,
			wantMessages: nil,
		},
		{
			path:         "testdata/0/code.go",
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
			path:         "testdata/2",
			includeTests: true,
			wantMessages: []string{
				"testdata/2/code.go:3 myVar is a global variable",
				"testdata/2/code_test.go:3 myTestVar is a global variable",
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
			path: "testdata/3/code_0.go",
			wantMessages: []string{
				"testdata/3/code_0.go:8 theVar is a global variable",
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
			path: "testdata/7",
			wantMessages: []string{
				"testdata/7/code.go:8 myVar is a global variable",
				// TODO: "testdata/7/code.go:13 errFakeErrorUnexported is a global variable",
				// TODO: "testdata/7/code.go:14 ErrFakeErrorExported is a global variable",
				"testdata/7/code.go:17 myErrVar is a global variable",
				"testdata/7/code.go:18 myVarErr is a global variable",
				"testdata/7/code.go:19 myVarError is a global variable",
				"testdata/7/code.go:20 customErr is a global variable",
				"testdata/7/code.go:29 declaredErr is a global variable",
			},
		},
		{
			path: "testdata/8",
			wantMessages: []string{
				"testdata/8/code.go:7 myVar is a global variable",
				// TODO: "testdata/8/code.go:7 errFakeErrorUnexported is a global variable",
				// TODO: "testdata/8/code.go:7 ErrFakeErrorExported is a global variable",
				"testdata/8/code.go:7 myErrVar is a global variable",
				"testdata/8/code.go:7 myVarErr is a global variable",
				"testdata/8/code.go:7 myVarError is a global variable",
				"testdata/8/code.go:7 customErr is a global variable",
				"testdata/8/code.go:7 declaredErr is a global variable",
			},
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
				"testdata/7/code.go:8 myVar is a global variable",
				// TODO: "testdata/7/code.go:13 errFakeErrorUnexported is a global variable",
				// TODO: "testdata/7/code.go:14 ErrFakeErrorExported is a global variable",
				"testdata/7/code.go:17 myErrVar is a global variable",
				"testdata/7/code.go:18 myVarErr is a global variable",
				"testdata/7/code.go:19 myVarError is a global variable",
				"testdata/7/code.go:20 customErr is a global variable",
				"testdata/7/code.go:29 declaredErr is a global variable",
				"testdata/8/code.go:7 myVar is a global variable",
				// TODO: "testdata/8/code.go:7 errFakeErrorUnexported is a global variable",
				// TODO: "testdata/8/code.go:7 ErrFakeErrorExported is a global variable",
				"testdata/8/code.go:7 myErrVar is a global variable",
				"testdata/8/code.go:7 myVarErr is a global variable",
				"testdata/8/code.go:7 myVarError is a global variable",
				"testdata/8/code.go:7 customErr is a global variable",
				"testdata/8/code.go:7 declaredErr is a global variable",
			},
		},
	}

	for _, c := range cases {
		caseName := fmt.Sprintf("%s include tests: %t", c.path, c.includeTests)
		t.Run(caseName, func(t *testing.T) {
			messages, err := checkNoGlobals(c.path, c.includeTests)
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
