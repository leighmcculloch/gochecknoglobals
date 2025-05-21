# gochecknoglobals

[![test](https://github.com/leighmcculloch/gochecknoglobals/actions/workflows/build.yml/badge.svg)](https://github.com/leighmcculloch/gochecknoglobals/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/gochecknoglobals)](https://goreportcard.com/report/github.com/leighmcculloch/gochecknoglobals)

Check that no globals are present in Go code.

## Why

Global variables are an input to functions that is not visible in the functions signature, complicate testing, reduces readability and increase the complexity of code.

https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html
https://twitter.com/davecheney/status/871939730761547776

### Linter Behavior and Mutation Detection Scope

The `gochecknoglobals` linter has evolved. Previously, it focused on flagging most global variable declarations with some exceptions. The current version shifts focus:
- **Global variable declarations are generally NOT flagged.**
- **Direct mutations of global variables ARE flagged.**

This approach allows for the existence of global variables (e.g., for configuration, singletons, error instances) but helps identify and prevent unintended or hard-to-track modifications to their state.

**Mutation Detection Details:**

1.  **Detected Mutations:** The linter currently detects direct assignments to global variables within function scopes. This includes:
    *   Simple assignment: `globalVar = newValue`
    *   Compound assignment: `globalVar += someValue`
    *   Reassignment of global pointers, slices, or maps: `globalSlice = newSlice`, `globalPtr = &newValue` (This means the global variable identifier itself is on the left-hand side of an assignment).

2.  **Undetected Mutations (Known Limitations):** The linter currently does **not** detect indirect mutations. These are modifications to the underlying data of a global variable without reassigning the global variable itself. Examples include:
    *   Modifying a field of a global struct: `globalStruct.Field = value`
    *   Modifying an element of a global array, slice, or map: `globalSlice[0] = value`, `globalMap["key"] = value`
    *   Mutations performed by functions or methods if a pointer to a global variable (or a global variable that is a pointer type like a slice or map) is passed to them: `modifyGlobal(&globalStruct)`, `globalInterface.Mutate()`.

3.  **Reason and Future Enhancements:** Detecting these indirect mutations is significantly more complex, requiring deeper data flow analysis. While not currently implemented, such enhancements could be considered for future development to provide more comprehensive state change tracking.

The original exception list for declarations (errors, version, regexp, go:embed) is no longer directly applicable for *allowing* declarations, as all declarations are implicitly allowed. However, the mutation of such variables will be flagged like any other global variable.

## Install

```
go install 4d63.com/gochecknoglobals@latest
```

## Usage

The linter is built on [Go's analysis package] and does thus support all the
built in flags and features from this type. The analyzer is executed by
specifying packages.

[Go's analysis package]: https://pkg.go.dev/golang.org/x/tools/go/analysis

```
gochecknoglobals [package]
```

```
gochecknoglobals ./...
```

By default, test files are checked but can be excluded by adding the
`-test=false` flag.

```
gochecknoglobals -test=false [package]
```
