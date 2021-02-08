package importer

import (
	"fmt"
	"sync"

	"github.com/fsilva1985/consultor-de-cep/console"
	"github.com/fsilva1985/consultor-de-cep/model"
	"gorm.io/gorm"
)

// Initialize instances
func Initialize(db *gorm.DB) {
	zipFile := Open("eDNE_Basico.zip")

	cities := ReadFile(
		zipFile,
		"Delimitado/LOG_LOCALIDADE.TXT",
	)

	City(cities, db)

	neighborhoods := ReadFile(
		zipFile,
		"Delimitado/LOG_BAIRRO.TXT",
	)

	Neighborhood(neighborhoods, db)

	var wg sync.WaitGroup

	wg.Add(len(model.GetStates()))
	for _, state := range model.GetStates() {
		file := ReadFile(
			zipFile,
			"Delimitado/LOG_LOGRADOURO_"+state.Code+".TXT",
		)

		Address(file, db, state.Code, &wg)
	}

	wg.Wait()

	fmt.Println(console.Messager("Estados importados com sucesso"))
}
