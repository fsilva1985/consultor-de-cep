package importer

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/fsilva1985/consultor-de-cep/console"
	"github.com/fsilva1985/consultor-de-cep/model"
	"github.com/fsilva1985/consultor-de-cep/parse"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// City returns void
func City(buffer *bufio.Scanner, db *gorm.DB) {
	var cities []model.City

	i := 1

	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		var state model.State

		db.First(&state, "code = ?", row[1])

		cities = append(cities, model.City{
			ID:      parse.StringToUint(row[0]),
			StateID: state.ID,
			Name:    row[2],
		})

		if i == 1000 {
			upsertCity(cities, db)
			i = 1
			cities = nil
		}

		i++
	}

	upsertCity(cities, db)

	fmt.Println(console.Messager("Cidades importados com sucesso"))
}

func upsertCity(data []model.City, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"state_id", "name"}),
	}).Create(data)
}
