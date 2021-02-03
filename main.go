package main

import (
	"github.com/fsilva1985/consultor-de-cep/archivezip"
	"github.com/fsilva1985/consultor-de-cep/importer"
	"github.com/fsilva1985/consultor-de-cep/model"
)

func main() {
	db := model.Init()

	zipFile := archivezip.Open("/storage/eDNE_Basico.zip")

	file := archivezip.ReadFile(
		zipFile,
		"Delimitado/LOG_LOCALIDADE.TXT",
	)

	importer.City(file, db)

	file = archivezip.ReadFile(
		zipFile,
		"Delimitado/LOG_BAIRRO.TXT",
	)

	importer.Neighborhood(file, db)

	for _, state := range model.GetStates() {
		file := archivezip.ReadFile(
			zipFile,
			"Delimitado/LOG_LOGRADOURO_"+state.Code+".TXT",
		)

		importer.Address(file, db, state.Code)
	}
}
