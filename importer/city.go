package importer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/djimenez/iconv-go"
	"github.com/fsilva1985/consultor-de-cep/console"
	"github.com/fsilva1985/consultor-de-cep/model"
	"github.com/fsilva1985/consultor-de-cep/parse"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// City returns void
func City(file io.ReadCloser, db *gorm.DB) {
	var cities []model.City

	i := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		stringer, _ := iconv.ConvertString(scanner.Text(), "windows-1252", "utf-8")

		row := strings.Split(stringer, "@")

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
