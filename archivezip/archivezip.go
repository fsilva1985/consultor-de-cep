package archivezip

import (
	"archive/zip"
	"bufio"
	"os"
	"strings"
)

// Open returns []byte
func Open(zipPath string) *zip.ReadCloser {
	pwd, _ := os.Getwd()

	read, _ := zip.OpenReader(pwd + zipPath)

	return read
}

// ReadFile returns *bufio.Scanner
func ReadFile(read *zip.ReadCloser, filepath string) *bufio.Scanner {
	var scanner *bufio.Scanner

	for _, file := range read.File {
		if strings.Compare(file.Name, filepath) != 0 {
			continue
		}

		buffer, _ := file.Open()

		scanner = bufio.NewScanner(buffer)
	}

	return scanner
}
