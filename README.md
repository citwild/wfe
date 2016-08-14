# Wide-Field Ethnography 

[![Build Status](https://travis-ci.org/citwild/wfe.svg?branch=master)](https://travis-ci.org/citwild/wfe)
[![Go Report Card](https://goreportcard.com/badge/github.com/citwild/wfe)](https://goreportcard.com/report/github.com/citwild/wfe)
[![GoDoc](https://godoc.org/github.com/citwild/wfe?status.svg)](https://godoc.org/github.com/citwild/wfe)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

[Wide-field ethnography](http://depts.washington.edu/citw/wordpress/?page_id=55) web app. Currently a work-in-progress.

## Requirements

- [Go](https://golang.org/doc/install)

Make sure you [set up your `GOPATH`](https://golang.org/doc/code.html#GOPATH). 

## Install

```
go get -u -v github.com/citwild/wfe/cmd/wfe
```

## Use

```
wfe -h
```

Or, if the `GOPATH`/bin directory is not in your `PATH`, on macOS and Linux:

```
"$GOPATH"/bin/wfe -h
```

On Windows:

```
"%GOPATH%"\bin\wfe -h
```

## License

The WFE web app is licensed under the [MIT License](https://opensource.org/licenses/MIT), which is an [open source license](https://opensource.org/docs/osd).
