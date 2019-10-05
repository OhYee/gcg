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
