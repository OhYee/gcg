package compare

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	type inputType struct {
		a Any
		b Any
	}
	type outputType struct {
		result int
		err    error
	}
	type TestCase struct {
		input  inputType
		output outputType
	}
	testCases := []TestCase{
		TestCase{inputType{1, 2}, outputType{-1, nil}},
		TestCase{inputType{2, 2}, outputType{0, nil}},
		TestCase{inputType{3, 2}, outputType{1, nil}},
		TestCase{inputType{int(1), int(2)}, outputType{-1, nil}},
		TestCase{inputType{uint8(1), uint8(2)}, outputType{-1, nil}},
		TestCase{inputType{uint16(1), uint16(2)}, outputType{-1, nil}},
		TestCase{inputType{uint32(1), uint32(2)}, outputType{-1, nil}},
		TestCase{inputType{uint64(1), uint64(2)}, outputType{-1, nil}},
		TestCase{inputType{float32(1), float32(2)}, outputType{-1, nil}},
		TestCase{inputType{float64(1), float64(2)}, outputType{-1, nil}},
		TestCase{inputType{1, uint8(2)}, outputType{-2, fmt.Errorf("Can not compare int and uint8")}},
	}
	for _, testCase := range testCases {
		result, err := Compare(testCase.input.a, testCase.input.b)
		if result != testCase.output.result || fmt.Sprintf("%v", err) != fmt.Sprintf("%v", testCase.output.err) {
			t.Errorf("Compare(%+v, %+v) = %+v, %+v, which expected %+v, %+v",
				testCase.input.a, testCase.input.b,
				result, err,
				testCase.output.result, testCase.output.err)
		} else {
			t.Logf("Compare(%+v, %+v) = %+v, %+v", testCase.input.a, testCase.input.b,
				result, err)
		}
	}
}
