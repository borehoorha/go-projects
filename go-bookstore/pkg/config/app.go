package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "borehoorha:0/simplerest?")
	if err != nil {
		panic(err)
	}
}
