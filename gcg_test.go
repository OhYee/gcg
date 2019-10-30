package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func writeToTempFile(t *testing.T, content string) string {
	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Errorf("Create temp file error: %v", err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		t.Errorf("Write to temp file error: %v", err)
	}
	file.Close()
	return file.Name()
}

func Test_renderTemplate(t *testing.T) {
	type testCase struct {
		name     string
		template string
		args     []interface{}
		result   string
	}
	tests := []testCase{
		{
			name:     "test upper",
			template: "{{upper .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "ABcdefg1234567",
				},
			},
			result: "ABCDEFG1234567",
		},
		{
			name:     "test lower",
			template: "{{lower .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "ABcdefg1234567",
				},
			},
			result: "abcdefg1234567",
		},
		{
			name:     "test upperFirstChar",
			template: "{{upperFirstChar .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "ABcdefg1234567",
				},
			},
			result: "ABcdefg1234567",
		},
		{
			name:     "test upperFirstChar 2",
			template: "{{upperFirstChar .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "abc",
				},
			},
			result: "Abc",
		},
		{
			name:     "test makeSlice",
			template: "{{$slice := makeSlice .a .b .c}}{{range $item := $slice}}{{$item}}{{end}}",
			args: []interface{}{
				map[string]interface{}{
					"a": "A", "b": "B", "c": "C",
				},
			},
			result: "ABC",
		},
		{
			name:     "test makeMap",
			template: "{{$map := makeMap .c .b .a  }}{{$map.C}}",
			args: []interface{}{
				map[string]interface{}{
					"a": "A", "b": "B", "c": "C",
				},
			},
			result: "B",
		},
	}

	martix := [7][9]interface{}{
		{"", "isInt", "isString", "isSlice", "isArray", "isMap", "isList", "isNumber", "isFloat"},
		{1, true, false, false, false, false, false, true, false},
		{"This is string", false, true, false, false, false, false, false, false},
		{make([]string, 3), false, false, true, false, false, true, false, false},
		{[2]string{"a", "b"}, false, false, false, true, false, true, false, false},
		{map[string]string{"A": "a"}, false, false, false, false, true, false, false, false},
		{0.7, false, false, false, false, false, false, true, true},
	}

	for i := 1; i < len(martix); i++ {
		value := martix[i][0]
		for j := 1; j < len(martix[i]); j++ {
			f := martix[0][j]
			r := "FALSE"
			if martix[i][j] == true {
				r = "TRUE"
			}
			tests = append(tests, testCase{
				name:     fmt.Sprintf("test %s with %v", f, value),
				template: fmt.Sprintf("{{if %s .}}TRUE{{else}}FALSE{{end}}", f),
				args:     []interface{}{value},
				result:   r,
			})
		}
	}

	buf := bytes.NewBuffer([]byte{})
	for _, tt := range tests {
		buf.Reset()
		filename := writeToTempFile(t, tt.template)
		t.Run(tt.name, func(t *testing.T) {
			renderTemplate(buf, []string{filename}, tt.args, map[string]interface{}{})
			if buf.String() != tt.result {
				t.Errorf("want %v, but got %v.", tt.result, buf.String())
			}
		})
		os.Remove(filename)
	}

}

func Test_modifyVariable(t *testing.T) {
	type args struct {
		args     interface{}
		variable map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "do nothing",
			args: args{
				args:     []interface{}{"a", "b", "c"},
				variable: map[string]interface{}{},
			},
			want: []interface{}{"a", "b", "c"},
		},
		{
			name: "modify direct string",
			args: args{
				args:     "$variable",
				variable: map[string]interface{}{"$variable": []interface{}{"a", "b", "c"}},
			},
			want: []interface{}{"a", "b", "c"},
		},
		{
			name: "modify in slice",
			args: args{
				args:     []interface{}{"$variable", "variable"},
				variable: map[string]interface{}{"$variable": []interface{}{"a", "b", "c"}},
			},
			want: []interface{}{[]interface{}{"a", "b", "c"}, "variable"},
		},
		{
			name: "modify in map",
			args: args{
				args:     map[string]interface{}{"key": "value", "variable": "$variable"},
				variable: map[string]interface{}{"$variable": []interface{}{"a", "b", "c"}},
			},
			want: map[string]interface{}{"key": "value", "variable": []interface{}{"a", "b", "c"}},
		},
		{
			name: "mix",
			args: args{
				args:     map[string]interface{}{"key": "value", "variable": []interface{}{"$variable", "variable"}},
				variable: map[string]interface{}{"$variable": []interface{}{"a", "b", "c"}},
			},
			want: map[string]interface{}{"key": "value", "variable": []interface{}{[]interface{}{"a", "b", "c"}, "variable"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifyVariable(tt.args.args, tt.args.variable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifyVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}
