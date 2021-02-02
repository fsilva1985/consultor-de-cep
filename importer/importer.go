package importer

import (
	"bufio"
	"strings"

	"github.com/fsilva1985/consultor-de-cep/model"
	"github.com/fsilva1985/consultor-de-cep/parse"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// City returns void
func City(buffer *bufio.Scanner, db *gorm.DB) {
	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"state_code", "name"}),
		}).Create(&model.City{
			Id:        parse.StringToUint(row[0]),
			StateCode: row[1],
			Name:      row[2],
		})
	}
}

// Neighborhood returns void
func Neighborhood(buffer *bufio.Scanner, db *gorm.DB) {
	for buffer.Scan() {
		row := strings.Split(buffer.Text(), "@")

		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"city_id", "name"}),
		}).Create(&model.Neighborhood{
			Id:     parse.StringToUint(row[0]),
			CityId: parse.StringToUint(row[2]),
			Name:   row[3],
		})
	}
}
