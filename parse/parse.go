package parse

import (
	"strconv"
)

// StringToUint returns uint
func StringToUint(value string) uint {
	number, _ := strconv.ParseUint(value, 10, 32)

	return uint(number)
}
