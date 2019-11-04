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

func Create(v interface{}) {
	db, err := gorm.Open("mysql", "root:idfcau1992@tcp(127.0.0.1:3306)/ams?charset=utf8&parseTime=True")
	checkError(err)
	defer db.Close()

	db.Create(v)
}
func UpdateByID(v interface{}, v1 interface{}) {
	db, err := gorm.Open("mysql", "root:idfcau1992@tcp(127.0.0.1:3306)/ams?charset=utf8&parseTime=True")
	checkError(err)
	defer db.Close()

	db.Find(v).Update(v1)
}
func MarkDelete(v interface{}) {
	db, err := gorm.Open("mysql", "root:idfcau1992@tcp(127.0.0.1:3306)/ams?charset=utf8&parseTime=True")
	checkError(err)
	defer db.Close()

	db.Find(v).Delete(v)
}

// func GetDB() *mgo.Database {
// 	return mgo.GetDB(cons.DB_ID)
// }

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
