# gochecknoglobals

[![test](https://github.com/leighmcculloch/gochecknoglobals/actions/workflows/build.yml/badge.svg)](https://github.com/leighmcculloch/gochecknoglobals/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/gochecknoglobals)](https://goreportcard.com/report/github.com/leighmcculloch/gochecknoglobals)

Check that no globals are present in Go code.

## Why

Global variables are an input to functions that is not visible in the functions signature, complicate testing, reduces readability and increase the complexity of code.

https://peter.bourgon.org/blog/2017/06/09/theory-of-modern-go.html
https://twitter.com/davecheney/status/871939730761547776

### Exceptions

There are very few exceptions to the global variable rule. This tool will ignore the following patterns:
 * Variables with an `Err` prefix
 * Variables named `_`
 * Variables named `version`
 * Variables assigned from `regexp.MustCompile()`
 * Variables with a `//go:embed` comment

## Install

```
go get 4d63.com/gochecknoglobals
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

By default, test files will not be checked but can be included by adding the
`-t` flag.

```
gochecknoglobals -t [package]
```
