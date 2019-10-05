# Go 代码生成器

![test state](https://github.com/OhYee/gcg/workflows/test/badge.svg)

[English](/README.md) | [中文](/README_CN.md)

使用 JSON 文件以及 Go 模板生成 Go 源码文件

## 安装

```
go get -u github.com/OhYee/gcg
go install github.com/OhYee/gcg
```

## 使用

```
gcg data.json
```

## 文档

JSON文件字段要求

|name|type||
|:---|:---|:---|
|variable|`map[string]interface{}`|全局变量|
|files|`[]goFile`|要生成的文件列表|

`goFile`结构如下:

|名称|类型||
|:---|:---|:---|
|package|string|生成文件的包名|
|import|`[]string` or `[][]string`|Go文件的引入包部分|
|output|`string`|输出文件名|
|body|`[]strut{template string, args []interface{}}` or `[]struct{template []string, args []interface{}}`|Go文件主体|

`body`的`template` 和 `args` 将用于 Go `text/template` 的参数

如果`args`的列表项为`string`,并且是`variable`的键，则它会被替换成对应的值

### 函数

|名称|参数||
|:---|:---|:---|
|lower|`string`|使所有字母小写|
|upper|`string`|使所有字母大写|
|upperFirstChar|`string`|使首字母大写|
|makeSlice|`[]interface{}`|将参数拼接成slice|
|makeMap|`[]interface{}`|将参数拼接成map|
|isInt|`interface{}`|判断参数是否为int|
|isString|`interface{}`|判断参数是否为string|
|isSlice|`interface{}`|判断参数是否为slice|
|isArray|`interface{}`|判断参数是否为array|
|isMap|`interface{}`|判断参数是否为map|
|isList|`interface{}`|判断参数是否为slice或array|
|isNumber|`interface{}`|判断参数是否为int或float|
|isFloat|`interface{}`|判断参数是否为float|

## 样例

详情见 [大小比较包](https://github.com/OhYee/gcg/tree/master/example/compare), 你可以先删除 `compare.go` 然后使用 `gcg data.json` 或者 `go generate g.go` 来重新生成它.

## 协议

[Apache License 2.0](/LICENSE)

