# Go 代码生成器

[English](/README.md) | [中文](/README_CN.md)

使用 JSON 文件以及 Go 模板生成 Go 源码文件

## 安装

```
go get github.com/OhYee/gcg
go install github.com/OhYee/gcg
```

## 使用

```
gcg data.json compare.go
```

## 文档

JSON文件字段要求

|名称|类型||
|:---|:---|:---|
|package|string|生成文件的包名|
|import|`[]string` or `[][]string`|Go文件的引入包部分|
|body|`[]strut{template string, args interface{}}` or `[]struct{template []string, args interface{}}`|Go文件主体|

`body`的`template` 和 `args` 将用于 Go `text/template` 的参数

### 函数

|名称|参数||
|:---|:---|:---|
|lower|`string`|使所有字母小写|
|upper|`string`|使所有字母大写|
|upperFirstChar|`string`|使首字母大写|

## 样例

详情见 [大小比较包](https://github.com/OhYee/gcg/tree/master/example/compare), 你可以先删除 `compare.go` 然后使用 `gcg data.json compare.go` 或者 `go generate g.go` 来重新生成它.

## 协议

[Apache License 2.0](/LICENSE)

