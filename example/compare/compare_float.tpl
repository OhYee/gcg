// {{.name}} compare {{.type}}
func {{.name}}(a {{.type}}, b {{.type}}) int {
	if math.Abs(float64(a-b)) < eps {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
