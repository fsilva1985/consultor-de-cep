package console

import (
	"runtime"
	"time"
)

// Messager returns string
func Messager(text string) string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + text
}
