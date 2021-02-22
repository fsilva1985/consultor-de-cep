package importers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/fsilva1985/consultor-de-cep/entities"
	"gorm.io/gorm"
)

// Initialize instances
func Initialize(db *gorm.DB) {
	zipFile := Open("eDNE_Basico.zip")

	cities := Read(
		zipFile,
		"Delimitado/LOG_LOCALIDADE.TXT",
	)

	City(cities, db)

	neighborhoods := Read(
		zipFile,
		"Delimitado/LOG_BAIRRO.TXT",
	)

	Neighborhood(neighborhoods, db)

	var wg sync.WaitGroup

	wg.Add(len(entities.GetStates()))
	for _, state := range entities.GetStates() {
		file := Read(
			zipFile,
			"Delimitado/LOG_LOGRADOURO_"+state.Code+".TXT",
		)

		Address(file, db, state.Code, &wg)
	}

	wg.Wait()

	fmt.Println("Estados importados com sucesso")
}

// Open returns []byte
func Open(zipPath string) *zip.ReadCloser {
	pwd, _ := os.Getwd()
	read, _ := zip.OpenReader(pwd + "/" + zipPath)

	return read
}

// ReadFile returns *bufio.Scanner
func Read(read *zip.ReadCloser, filepath string) io.ReadCloser {
	var file io.ReadCloser

	for _, currentFile := range read.File {
		if strings.Compare(currentFile.Name, filepath) != 0 {
			continue
		}

		file, _ = currentFile.Open()
	}

	return file
}
