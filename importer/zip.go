package importer

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// Open returns []byte
func Open(zipPath string) *zip.ReadCloser {
	pwd, _ := os.Getwd()
	read, _ := zip.OpenReader(pwd + "/" + zipPath)

	return read
}

// ReadFile returns *bufio.Scanner
func ReadFile(read *zip.ReadCloser, filepath string) io.ReadCloser {
	var file io.ReadCloser

	for _, currentFile := range read.File {
		if strings.Compare(currentFile.Name, filepath) != 0 {
			continue
		}

		file, _ = currentFile.Open()
	}

	return file
}
