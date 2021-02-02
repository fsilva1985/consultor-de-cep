package main

import (
	"fmt"
	"time"

	"github.com/fsilva1985/consultor-de-cep/importer"
	"github.com/fsilva1985/consultor-de-cep/model"
	"github.com/fsilva1985/consultor-de-cep/zip"
)

func main() {
	db := model.Init()

	zipFile := zip.Open("/storage/eDNE_Basico.zip")

	file := zip.ReadFile(
		zipFile,
		"Delimitado/LOG_LOCALIDADE.TXT",
	)

	fmt.Println("Cidades Inicio", time.Now().Format("2006-01-02 15:04:05"))
	importer.City(file, db)
	fmt.Println("Cidades Fim", time.Now().Format("2006-01-02 15:04:05"))

	file = zip.ReadFile(
		zipFile,
		"Delimitado/LOG_BAIRRO.TXT",
	)

	fmt.Println("Bairros Inicio", time.Now().Format("2006-01-02 15:04:05"))
	importer.Neighborhood(file, db)
	fmt.Println("Bairros Fim", time.Now().Format("2006-01-02 15:04:05"))
}
