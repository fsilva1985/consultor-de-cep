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

	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		var state model.State

		db.First(&state, "code = ?", row[1])

		cities = append(cities, model.City{
			Id:      parse.StringToUint(row[0]),
			StateId: state.Id,
			Name:    row[2],
		})
	}

	var slices [][]model.City = createCityChunk(cities, 100)

	done := make(chan string)
	go upsertCity(slices, db, done)
	fmt.Println(<-done)
}

func upsertCity(slices [][]model.City, db *gorm.DB, done chan string) {
	for _, slice := range slices {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"state_id", "name"}),
		}).Create(slice)
	}
	done <- console.GetDoneMessage("Cidades importados")
}

// createCityChunk returns [][]model.City
func createCityChunk(arr []model.City, limit int) [][]model.City {
	var slices [][]model.City
	for i := 0; i < len(arr); i += limit {
		slices = append(slices, arr[i:i+limit])
	}

	return slices
}
