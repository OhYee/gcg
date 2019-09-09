# Go Code Generator

[English](/README.md) | [中文](/README_CN.md)

Using json file and go template to generate go code.

## Installation

```bash
go get github.com/OhYee/gcg
go install github.com/OhYee/gcg
```

## Usage

```bash
gcg data.json compare.go
```

## Document

The json file must have these field:

|name|type||
|:---|:---|:---|
|package|string|package name|
|import|`[]string` or `[][]string`|import part of the file|
|body|`[]strut{template string, args interface{}}` or `[]struct{template []string, args interface{}}`|the go file body|

The `template` and `args` in `body` will be the arguments of go `text/template`

### Function

|name|args||
|:---|:---|:---|
|lower|`string`|make all letters to their lower case|
|upper|`string`|make all letters to their upper case|
|upperFirstChar|`string`|make the first letter to its upper case, keep others|

## Example

See [number compare package](https://github.com/OhYee/gcg/tree/master/example/compare), you can delete `compare.go` and using `gcg data.json compare.go` or `go generate g.go` to re-generate it.

## LICENSE

[Apache License 2.0](/LICENSE)

