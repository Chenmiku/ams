package database

import (
	"ams/dapi/config/cons"
	"db/mgo"
	"fmt"
)

type DatabaseConfig struct {
	DBHost   string
	DBName   string
	UserName string
	PassWord string
}

func (o DatabaseConfig) String() string {
	return fmt.Sprintf("db:host=%s;name=%s", o.DBHost, o.DBName)
}

func (o *DatabaseConfig) Check() {
	mgo.Register(cons.DB_ID, o.DBName, o.DBHost, o.UserName, o.PassWord)
}
