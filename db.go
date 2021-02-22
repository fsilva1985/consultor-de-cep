package main

import (
	"github.com/fsilva1985/consultor-de-cep/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getDBInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.State{}, &entities.City{}, &entities.Neighborhood{}, &entities.Address{})

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"code", "name"}),
	}).Create(entities.GetStates())

	return db
}
