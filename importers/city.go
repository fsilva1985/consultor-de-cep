package importers

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/djimenez/iconv-go"
	"github.com/fsilva1985/consultor-de-cep/entities"
	"github.com/fsilva1985/consultor-de-cep/parsers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// City returns void
func City(file io.ReadCloser, db *gorm.DB) {
	var cities []entities.City

	i := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		stringer, _ := iconv.ConvertString(scanner.Text(), "windows-1252", "utf-8")

		row := strings.Split(stringer, "@")

		var state entities.State

		db.First(&state, "code = ?", row[1])

		cities = append(cities, entities.City{
			ID:      parsers.StringToUint(row[0]),
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

	fmt.Println("Cidades importados com sucesso")
}

func upsertCity(data []entities.City, db *gorm.DB) {
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"state_id", "name"}),
	}).Create(data)
}
