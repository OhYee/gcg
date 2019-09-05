package compare

import (
	"fmt"
	"math" // math.Abs()
)

// Any random type variable
type Any = interface{}

const eps = 1e-7

// Compare interface{} a and b
//
// a > b: 1
// a = b: 0
// a < b: -1
// if a and b are different type, then result = -2 and return error
func Compare(a Any, b Any) (int, error) {
	compareError := fmt.Errorf("Can not compare %T and %T", a, b)

	switch a.(type) {
	case int:
		dataB, ok := b.(int)
		if ok {
			return Int(a.(int), dataB), nil
		}
		return -2, compareError
	case uint8:
		dataB, ok := b.(uint8)
		if ok {
			return Uint8(a.(uint8), dataB), nil
		}
		return -2, compareError
	case uint16:
		dataB, ok := b.(uint16)
		if ok {
			return Uint16(a.(uint16), dataB), nil
		}
		return -2, compareError
	case uint32:
		dataB, ok := b.(uint32)
		if ok {
			return Uint32(a.(uint32), dataB), nil
		}
		return -2, compareError
	case uint64:
		dataB, ok := b.(uint64)
		if ok {
			return Uint64(a.(uint64), dataB), nil
		}
		return -2, compareError
	case float32:
		dataB, ok := b.(float32)
		if ok {
			return Float32(a.(float32), dataB), nil
		}
		return -2, compareError
	case float64:
		dataB, ok := b.(float64)
		if ok {
			return Float64(a.(float64), dataB), nil
		}
		return -2, compareError
	}
	return -2, compareError
}

// Int compare int
func Int(a int, b int) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Uint8 compare uint8
func Uint8(a uint8, b uint8) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Uint16 compare uint16
func Uint16(a uint16, b uint16) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Uint32 compare uint32
func Uint32(a uint32, b uint32) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Uint64 compare uint64
func Uint64(a uint64, b uint64) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Int8 compare int8
func Int8(a int8, b int8) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Int16 compare int16
func Int16(a int16, b int16) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Int32 compare int32
func Int32(a int32, b int32) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Int64 compare int64
func Int64(a int64, b int64) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	case a > b:
		return 1
	}
	return 0
}

// Float32 compare float32
func Float32(a float32, b float32) int {
	if math.Abs(float64(a-b)) < eps {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

// Float64 compare float64
func Float64(a float64, b float64) int {
	if math.Abs(float64(a-b)) < eps {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
