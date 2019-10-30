package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"gopkg.in/ffmt.v1"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"text/template"
)

const (
	version  = "0.0.6"
	helpText = "Using `gcg <json file>` to generate go file\nSuch as `gcg data.json`"
)

// CGO_ENABLED: 0
// GOOS: darwin、freebsd、linux、windows
// GOARCH: 386、amd64、arm
//
//go:generate bash -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gcg gcg.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/gcg.exe gcg.go"
//go:generate bash -c "CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/gcg_x86.exe gcg.go"

type config struct {
	Variable map[string]interface{} `json:"variable"`
	Files    []goFile               `json:"files"`
	Root     string
}

type goFile struct {
	PackageName     string        `json:"package"`
	ImportedPackage []interface{} `json:"import"`
	Body            []bodyArea    `json:"body"`
	Output          string        `json:"output"`
}

type bodyArea struct {
	Template interface{}   `json:"template"`
	Args     []interface{} `json:"args"`
}
type jsonMap = map[string]interface{}

var funcMap = template.FuncMap{
	"upperFirstChar": func(text string) string {
		return fmt.Sprintf("%s%s", strings.ToUpper(text[0:1]), text[1:len(text)])
	},
	"upper": func(text string) string {
		return strings.ToUpper(text)
	},
	"lower": func(text string) string {
		return strings.ToLower(text)
	},
	"makeSlice": func(args ...interface{}) (slice []interface{}) {
		for _, arg := range args {
			slice = append(slice, arg)
		}
		return
	},
	"makeMap": func(args ...interface{}) (m map[interface{}]interface{}) {
		m = make(map[interface{}]interface{})
		for i := 0; i < len(args)/2; i++ {
			m[args[i*2]] = args[i*2+1]
		}
		return
	},
	"isInt": func(i interface{}) bool {
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
			return true
		default:
			return false
		}
	},
	"isString": func(i interface{}) bool {
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.String:
			return true
		default:
			return false
		}
	},
	"isSlice": func(i interface{}) bool {
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Slice:
			return true
		default:
			return false
		}
	},
	"isArray": func(i interface{}) bool {
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Array:
			return true
		default:
			return false
		}
	},
	"isMap": func(i interface{}) bool {
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Map:
			return true
		default:
			return false
		}
	},
	"isList": func(i interface{}) bool {
		v := reflect.ValueOf(i).Kind()
		return v == reflect.Array || v == reflect.Slice
	},
	"isNumber": func(i interface{}) bool {
		v := reflect.ValueOf(i).Kind()
		switch v {
		case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			return true
		default:
			return false
		}
	},
	"isFloat": func(i interface{}) bool {
		v := reflect.ValueOf(i).Kind()
		return v == reflect.Float32 || v == reflect.Float64
	},
}

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
func readData(filename string) (cfg config) {
	var err error

	rawText, err := ioutil.ReadFile(filename)
	exitWhenError(err)

	err = json.Unmarshal(rawText, &cfg)
	exitWhenError(err)

	filenameSlice := strings.Split(filename, "/")
	cfg.Root = strings.Join(filenameSlice[0:len(filenameSlice)-1], "/")
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

// getFilename from a path (go template name can not be a path)
func getFileName(path string) (filename string) {
	pathList := strings.Split(path, "/")
	filename = pathList[len(pathList)-1]
	return
}

// renderTemplate render the block template
func renderTemplate(buf io.Writer, templates []string, args []interface{}, variable map[string]interface{}) {
	templateName := getFileName(templates[0])
	tpl, err := template.New(templateName).Funcs(funcMap).ParseFiles(templates...)
	exitWhenError(err)
	for _, arg := range args {
		// tpl.Execute(buf, arg)

		var temp interface{}
		s, ok := arg.(string)
		if ok {
			temp, ok = variable[s]
		}
		if !ok {
			temp = arg
		}

		err = tpl.ExecuteTemplate(buf, templateName, temp)
		exitWhenError(err)
		// buf.Write([]byte{'\n'})
	}
}

// renderContent render generated file content
func renderContent(cfg config, gf goFile) {
	buf := bytes.NewBufferString(fmt.Sprintf("// Code generated by GCG. DO NOT EDIT.\n// Go Code Generator %v (https://github.com/OhYee/gcg)\n\n", version))
	buf.WriteString(fmt.Sprintf("package %s\n\n", gf.PackageName))
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
		}(gf.ImportedPackage),
	))

	for _, block := range gf.Body {
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
		for idx := range templates {
			templates[idx] = fmt.Sprintf("%s/%s", cfg.Root, templates[idx])
		}

		renderTemplate(buf, templates, block.Args, cfg.Variable)

	}
	// code format
	b, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("%s: goformat error\n%s\n", gf.Output, err)
		b = buf.Bytes()
	}

	outputFile, err := os.Create(fmt.Sprintf("%s/%s", cfg.Root, gf.Output))
	exitWhenError(err)
	outputFile.Write(b)
	outputFile.Close()
	return
}

func main() {
	var inputFile string
	switch len(os.Args) {
	case 1:
		exitWhenFalse(false, helpText)
	default:
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			exitWhenFalse(false, helpText)
		} else if os.Args[1] == "-v" || os.Args[1] == "--version" {
			exitWhenFalse(false, fmt.Sprintf("Go Code Generator version: %s\n", version))
		} else if os.Args[1] == "-d" || os.Args[1] == "--debug" {
			ffmt.P(readData(os.Args[2]))
			exitWhenFalse(false, "\n")
		}
		inputFile = os.Args[1]

	}
	cfg := readData(inputFile)
	for _, gf := range cfg.Files {
		renderContent(cfg, gf)
	}
}
