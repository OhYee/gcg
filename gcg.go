package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

type arguments struct {
	PackageName     string        `json:"package"`
	ImportedPackage []interface{} `json:"import"`
	Body            []bodyArea    `json:"body"`
}
type bodyArea struct {
	Template interface{}   `json:"template"`
	Args     []interface{} `json:"args"`
}
type jsonMap = map[string]interface{}

// exitWhenError exit the program when error is not null
func exitWhenError(err error) {
	if err != nil {
		fmt.Printf("Can not solve json data: %v", err)
		os.Exit(0)
	}
}

// exitWhenFalse exit the program when value is false
func exitWhenFalse(b bool, errMsg ...interface{}) {
	if !b {
		fmt.Println(errMsg...)
		os.Exit(0)
	}
}

// readData read arguments from json
func readData(filename string) (args arguments) {
	var err error

	rawText, err := ioutil.ReadFile(filename)
	exitWhenError(err)

	err = json.Unmarshal(rawText, &args)
	exitWhenError(err)

	return
}

// importPackage from a string or []string to generator import part
func importPackage(x interface{}) (s string) {
	switch x.(type) {
	case string:
		str, _ := x.(string)
		s = fmt.Sprintf("\"%s\"", str)
	case []interface{}:
		slice, _ := x.([]interface{})
		switch len(slice) {
		case 0:
			s = ""
		case 1:
			s = fmt.Sprintf("\"%s\"", slice[0])
		case 2:
			s = fmt.Sprintf("%s \"%s\"", slice[0], slice[1])
		default:
			s = fmt.Sprintf("%s \"%s\" // %s", slice[0], slice[1], slice[2])
		}
	default:
		exitWhenError(fmt.Errorf("Type of %v is %T, expected string or []interface{}", x, x))
	}
	return
}

func main() {
	var inputFile string
	var outputFile *os.File
	switch len(os.Args) {
	case 1:
		exitWhenFalse(false, "Using `gcg <json file> [<output file>]` to generate go file\nSuch as `gcg data.json` or `gcg data.json ../add.go`")
	case 2:
		inputFile = os.Args[1]
		outputFile = os.Stdout
	default:
		var err error
		inputFile = os.Args[1]
		outputFile, err = os.Create(os.Args[2])
		exitWhenError(err)
	}
	args := readData(inputFile)

	buf := bytes.NewBufferString(fmt.Sprintf("package %s\n\n", args.PackageName))
	buf.WriteString(fmt.Sprintf("import (\n%s\n)\n\n",
		func(slice []interface{}) string {
			s := ""
			for idx, pkg := range slice {
				s += fmt.Sprintf("%s\t%s", func(idx int) string {
					if idx > 0 {
						return "\n"
					}
					return ""
				}(idx), importPackage(pkg))
			}
			return s
		}(args.ImportedPackage),
	))

	for _, block := range args.Body {
		var templates []string
		switch block.Template.(type) {
		case string:
			temp, _ := block.Template.(string)
			templates = []string{temp}
		case []interface{}:
			var isString bool
			temp, _ := block.Template.([]interface{})
			templates = make([]string, len(temp))
			for idx, item := range temp {
				templates[idx], isString = item.(string)
				exitWhenFalse(isString, "template must be string or []string")
			}
		default:
			exitWhenFalse(false, "template must be string or []string")
		}

		tpl, err := template.ParseFiles(templates...)
		exitWhenError(err)
		for _, arg := range block.Args {
			tpl.Execute(buf, arg)
			buf.WriteString("\n\n")
		}
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("goformat error\n%s\n", err)
		b = buf.Bytes()
	}
	outputFile.Write(b)
}
