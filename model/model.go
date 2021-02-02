package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type City struct {
	Id        uint
	StateCode string
	Name      string
}

type Neighborhood struct {
	Id     uint
	CityId uint
	Name   string
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("model/database.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&City{}, &Neighborhood{})

	return db
}
