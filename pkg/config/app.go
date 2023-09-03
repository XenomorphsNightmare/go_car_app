package config

import (
	"github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db * gorm.DB
)

func Connect () {
	d, err := gorm.Open("sqlite3","test.db" )
	if err != nil {
		panic("Can't reach database")
	}

	db = d 
}

func GetDB() *gorm.DB{
	return db
}