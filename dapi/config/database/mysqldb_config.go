package database

import (
	// "ams/dapi/config/cons"
	"fmt"
	"github.com/jinzhu/gorm"
)

type DatabaseConfig struct {
	Driver   string
	DBHost   string
	DBName   string
	UserName string
	PassWord string
}

func (o DatabaseConfig) String() string {
	return fmt.Sprintf("db:host=%s;name=%s", o.DBHost, o.DBName)
}

func (o *DatabaseConfig) Check() {
	db, err := gorm.Open(o.Driver, o.UserName+":"+o.PassWord+"@tcp("+o.DBHost+")/"+o.DBName+"?charset=utf8&parseTime=True")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
