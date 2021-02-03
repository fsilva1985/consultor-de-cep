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

// Neighborhood returns void
func Neighborhood(buffer *bufio.Scanner, db *gorm.DB) {
	var neighborhoods []model.Neighborhood
	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		neighborhoods = append(neighborhoods, model.Neighborhood{
			Id:     parse.StringToUint(row[0]),
			CityId: parse.StringToUint(row[2]),
			Name:   row[3],
		})

	}

	var slices [][]model.Neighborhood = createNeighborhoodChunk(neighborhoods, 100)

	done := make(chan string)
	go upsertNeighborhood(slices, db, done)
	fmt.Println(<-done)
}

func upsertNeighborhood(slices [][]model.Neighborhood, db *gorm.DB, done chan string) {
	for _, slice := range slices {
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"city_id", "name"}),
		}).Create(slice)
	}
	done <- console.GetDoneMessage("Bairros importados")
}

// createNeighborhoodChunk returns [][]model.Neighborhood
func createNeighborhoodChunk(arr []model.Neighborhood, limit int) [][]model.Neighborhood {
	var slices [][]model.Neighborhood
	for i := 0; i < len(arr); i += limit {
		slices = append(slices, arr[i:i+limit])
	}

	return slices
}
