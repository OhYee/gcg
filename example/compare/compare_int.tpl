// {{.name}} compare {{.type}}
func {{.name}}(a {{.type}}, b {{.type}}) int {
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
