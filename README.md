# Go Code Generator

[English](/README.md) | [中文](/README_CN.md)

Using json file and go template to generate go code.

## Installation

```
go get github.com/OhYee/gcg
go install github.com/OhYee/gcg
```

## Usage

```
gcg data.json compare.go
```

## document

The json file must have these field:

|name|type||
|:---:|:---:|:---:|
|package|string|package name|
|import|`[]string` or `[][]string`|import part of the file|
|body|`[]strut{template string, args interface{}}` or `[]struct{template []string, args interface{}}`|the go file body|

The `template` and `args` in `body` will be the arguments of go `text/template`

## Example

See [number compare package](https://github.com/OhYee/gcg/tree/master/example/compare), you can delete `compare.go` and using `gcg data.json compare.go` or `go generate g.go` to re-generate it.

## LICENSE

[Apache License 2.0](/LICENSE)

