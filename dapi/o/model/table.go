package model

import (
	// "ams/dapi/config/cons"
	"github.com/jinzhu/gorm"
)

func NewTable(v interface{}) {
	db, err := gorm.Open("mysql", "root:idfcau1992@tcp(127.0.0.1:3306)/ams?charset=utf8&parseTime=True")
	checkError(err)
	defer db.Close()

	db.AutoMigrate(v)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
