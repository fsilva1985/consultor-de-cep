package importer

import (
	"github.com/fsilva1985/consultor-de-cep/model"
	"gorm.io/gorm"
)

// Initialize instances
func Initialize(db *gorm.DB) {
	zipFile := Open("eDNE_Basico.zip")

	file := ReadFile(
		zipFile,
		"Delimitado/LOG_LOCALIDADE.TXT",
	)

	City(file, db)

	file = ReadFile(
		zipFile,
		"Delimitado/LOG_BAIRRO.TXT",
	)

	Neighborhood(file, db)

	file = ReadFile(
		zipFile,
		"Delimitado/LOG_BAIRRO.TXT",
	)

	for _, state := range model.GetStates() {
		file := ReadFile(
			zipFile,
			"Delimitado/LOG_LOGRADOURO_"+state.Code+".TXT",
		)

		Address(file, db, state.Code)
	}
}
