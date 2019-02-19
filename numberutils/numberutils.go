package numberutils

import "strconv"

// ItfToInt interface{} to int
func ItfToInt(i64 interface{}) (i int) {
	switch v := i64.(type) {
	case int:
		i = v
	case int32:
		i = int(v)
	case int64:
		i = int(v)
	case float32:
		i = int(v)
	case float64:
		i = int(v)
	case string:
		vF64, _ := strconv.ParseFloat(v, 64)
		i = int(vF64)
	default:
		i = 0
	}
	return
}
