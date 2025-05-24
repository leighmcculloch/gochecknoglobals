# gochecknoglobals

[![test](https://github.com/leighmcculloch/gochecknoglobals/actions/workflows/build.yml/badge.svg)](https://github.com/leighmcculloch/gochecknoglobals/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/gochecknoglobals)](https://goreportcard.com/report/github.com/leighmcculloch/gochecknoglobals)

Check that global variables are not mutated in Go code.

## Why

Global variables are an input to functions that is not visible in the functions signature, complicate testing, reduces readability and increase the complexity of code. Side-effects arise when global variables are mutated because any function in a package can change an unexported package variable, and any function anywhere in an application can change an exported package variable.

This linter allows global variables but disallows mutation of them after initialization, encouraging them to be used like constants.

https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html
https://twitter.com/davecheney/status/871939730761547776

### Exceptions

Previously, this tool would error on all global variables with a few exceptions. 
Now it allows global variables but errors on mutations of them, encouraging globals 
to be used as pseudo-constants.

The linter will report any assignment, increment, or decrement of a global variable 
that occurs outside of the variable's initialization.

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
