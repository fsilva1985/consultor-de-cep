package console

import (
	"runtime"
	"strconv"
	"time"
)

// GetDoneMessage returns string
func GetDoneMessage(text string) string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + text + " | MemoryUsage " + bToMb(m.Alloc) + " MB"
}

func bToMb(b uint64) string {
	integer := int(b / 1024 / 1024)

	return strconv.Itoa(integer)
}
