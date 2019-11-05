package database

import (
	"ams/dapi/o/auth/session"
	"ams/dapi/o/org/role"
	"ams/dapi/o/org/user"
	"log"
	// "ams/dapi/config/cons"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DatabaseConfig struct {
	Driver   string
	DBHost   string
	DBName   string
	UserName string
	PassWord string
	db       *gorm.DB
}

func (o DatabaseConfig) String() string {
	return fmt.Sprintf("db:host=%s;name=%s", o.DBHost, o.DBName)
}

func (o *DatabaseConfig) Check() {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", o.UserName, o.PassWord, o.DBHost, o.DBName)
	o.db, err = gorm.Open(o.Driver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database ", o.DBName)
		log.Fatal("Error connection: ", err)
	} else {
		fmt.Printf("We are connected to the %s database ", o.DBName)
	}

	o.db.Debug().AutoMigrate(&user.User{}, &session.Session{}, &role.Role{}) //database migration

	defer o.db.Close()
}
