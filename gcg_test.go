package main

import (
	"bytes"
	"io/ioutil"
	"os"
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
	tests := []struct {
		name     string
		template string
		args     []interface{}
		result   string
	}{
		{
			name:     "test upper",
			template: "{{upper .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "ABcdefg1234567",
				},
			},
			result: "ABCDEFG1234567\n",
		},
		{
			name:     "test lower",
			template: "{{lower .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "ABcdefg1234567",
				},
			},
			result: "abcdefg1234567\n",
		},
		{
			name:     "test upperFirstChar",
			template: "{{upperFirstChar .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "ABcdefg1234567",
				},
			},
			result: "ABcdefg1234567\n",
		},
		{
			name:     "test upperFirstChar 2",
			template: "{{upperFirstChar .text}}",
			args: []interface{}{
				map[string]interface{}{
					"text": "abc",
				},
			},
			result: "Abc\n",
		},
	}

	buf := bytes.NewBuffer([]byte{})
	for _, tt := range tests {
		buf.Reset()
		filename := writeToTempFile(t, tt.template)
		t.Run(tt.name, func(t *testing.T) {
			renderTemplate(buf, []string{filename}, tt.args)
			if buf.String() != tt.result {
				t.Errorf("want %v, but got %v", tt.result, buf.String())
			}
		})
		os.Remove(filename)
	}
}
