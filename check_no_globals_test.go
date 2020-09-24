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
				"testdata/7/code.go:13 errFakeErrorUnexported is a global variable",
				"testdata/7/code.go:14 ErrFakeErrorExported is a global variable",
				"testdata/7/code.go:17 myErrVar is a global variable",
				"testdata/7/code.go:18 myVarErr is a global variable",
				"testdata/7/code.go:19 myVarError is a global variable",
				"testdata/7/code.go:20 customErr is a global variable",
				"testdata/7/code.go:25 errCustomUnexported is a global variable",
				"testdata/7/code.go:26 ErrCustomExported is a global variable",
				"testdata/7/code.go:29 declaredErr is a global variable",
			},
		},
		{
			path: "testdata/8",
			wantMessages: []string{
				"testdata/8/code.go:9 myVar is a global variable",
				"testdata/8/code.go:14 errFakeErrorUnexported is a global variable",
				"testdata/8/code.go:15 ErrFakeErrorExported is a global variable",
				"testdata/8/code.go:18 myErrVar is a global variable",
				"testdata/8/code.go:19 myVarErr is a global variable",
				"testdata/8/code.go:20 myVarError is a global variable",
				"testdata/8/code.go:21 customErr is a global variable",
				"testdata/8/code.go:26 errCustomUnexported is a global variable",
				"testdata/8/code.go:27 ErrCustomExported is a global variable",
				"testdata/8/code.go:30 declaredErr is a global variable",
			},
		},
		{
			path: "testdata/9",
			wantMessages: []string{
				"testdata/9/code.go:3 Version is a global variable",
				"testdata/9/code.go:4 version22 is a global variable",
				"testdata/9/code.go:5 version is a global variable",
			},
		},
		{
			path: "testdata/11",
			wantMessages: []string{
				"testdata/11/code.go:9 allWrong is a global variable",
				"testdata/11/code.go:12 nameWrong is a global variable",
				"testdata/11/code.go:13 nameWrongAgain is a global variable",
				"testdata/11/code.go:16 ErrCorrect is a global variable",
				"testdata/11/code.go:17 ErrAlsoCorrect is a global variable",
				"testdata/11/code.go:25 allWrongGroup is a global variable",
				"testdata/11/code.go:26 nameWrongGroup is a global variable",
				"testdata/11/code.go:27 nameWrongAgainGroup is a global variable",
				"testdata/11/code.go:28 ErrCorrectGroup is a global variable",
				"testdata/11/code.go:29 ErrAlsoCorrectGroup is a global variable",
			},
		},
		{
			path:         ".",
			wantMessages: nil,
		},
		{
			path: "./...",
			wantMessages: []string{
				"testdata/10/code.go:10 myVar is a global variable",
				"testdata/10/code.go:16 ErrIsNotErr is a global variable",
				"testdata/10/code.go:33 HTTPClient is a global variable",
				"testdata/11/code.go:9 allWrong is a global variable",
				"testdata/11/code.go:12 nameWrong is a global variable",
				"testdata/11/code.go:13 nameWrongAgain is a global variable",
				"testdata/11/code.go:16 ErrCorrect is a global variable",
				"testdata/11/code.go:17 ErrAlsoCorrect is a global variable",
				"testdata/11/code.go:25 allWrongGroup is a global variable",
				"testdata/11/code.go:26 nameWrongGroup is a global variable",
				"testdata/11/code.go:27 nameWrongAgainGroup is a global variable",
				"testdata/11/code.go:28 ErrCorrectGroup is a global variable",
				"testdata/11/code.go:29 ErrAlsoCorrectGroup is a global variable",
				"testdata/2/code.go:3 myVar is a global variable",
				"testdata/3/code_0.go:8 theVar is a global variable",
				"testdata/3/code_1.go:3 myVar is a global variable",
				"testdata/4/code.go:3 theVar is a global variable",
				"testdata/4/subpkg/code_1.go:3 myVar is a global variable",
				"testdata/5/code.go:3 myVar1 is a global variable",
				"testdata/5/code.go:3 myVar2 is a global variable",
				"testdata/7/code.go:8 myVar is a global variable",
				"testdata/7/code.go:13 errFakeErrorUnexported is a global variable",
				"testdata/7/code.go:14 ErrFakeErrorExported is a global variable",
				"testdata/7/code.go:17 myErrVar is a global variable",
				"testdata/7/code.go:18 myVarErr is a global variable",
				"testdata/7/code.go:19 myVarError is a global variable",
				"testdata/7/code.go:20 customErr is a global variable",
				"testdata/7/code.go:25 errCustomUnexported is a global variable",
				"testdata/7/code.go:26 ErrCustomExported is a global variable",
				"testdata/7/code.go:29 declaredErr is a global variable",
				"testdata/8/code.go:9 myVar is a global variable",
				"testdata/8/code.go:14 errFakeErrorUnexported is a global variable",
				"testdata/8/code.go:15 ErrFakeErrorExported is a global variable",
				"testdata/8/code.go:18 myErrVar is a global variable",
				"testdata/8/code.go:19 myVarErr is a global variable",
				"testdata/8/code.go:20 myVarError is a global variable",
				"testdata/8/code.go:21 customErr is a global variable",
				"testdata/8/code.go:26 errCustomUnexported is a global variable",
				"testdata/8/code.go:27 ErrCustomExported is a global variable",
				"testdata/8/code.go:30 declaredErr is a global variable",
				"testdata/9/code.go:3 Version is a global variable",
				"testdata/9/code.go:4 version22 is a global variable",
				"testdata/9/code.go:5 version is a global variable",
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
