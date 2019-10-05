# Go Code Generator

![test state](https://github.com/OhYee/gcg/workflows/test/badge.svg)

[English](/README.md) | [中文](/README_CN.md)

Using json file and go template to generate go code.

## Installation

```bash
go get -u github.com/OhYee/gcg
go install github.com/OhYee/gcg
```

## Usage

```bash
gcg data.json
```

## Document

The json file must have these field:


|name|type||
|:---|:---|:---|
|variable|`map[string]interface{}`|gloable variable|
|files|`[]goFile`|the list of generated files|

goFile struct like below:

|name|type||
|:---|:---|:---|
|package|`string`|package name|
|import|`[]string` or `[][]string`|import part of the file|
|output|`string`|the output filename|
|body|`[]strut{template string, args []interface{}}` or `[]struct{template []string, args []interface{}}`|the go file body|

The `template` and `args` in `body` will be the arguments of go `text/template`

If the item of `args` is `string`, and match `variable`'s key, it will be replaced to the value.

### Function

|name|args||
|:---|:---|:---|
|lower|`string`|make all letters to their lower case|
|upper|`string`|make all letters to their upper case|
|upperFirstChar|`string`|make the first letter to its upper case, keep others|
|makeSlice|`[]interface{}`|concat arguments as slice|
|makeMap|`[]interface{}`|concat arguments as map|
|isInt|`interface{}`|judge argument is int|
|isString|`interface{}`|judge argument is string|
|isSlice|`interface{}`|judge argument is slice|
|isArray|`interface{}`|judge argument is array|
|isMap|`interface{}`|judge argument is map|
|isList|`interface{}`|judge argument is slice or array|
|isNumber|`interface{}`|judge argument is int or float|
|isFloat|`interface{}`|judge argument is float|

## Example

See [number compare package](https://github.com/OhYee/gcg/tree/master/example/compare), you can delete `compare.go` and using `gcg data.json` or `go generate g.go` to re-generate it.

## LICENSE

[Apache License 2.0](/LICENSE)

