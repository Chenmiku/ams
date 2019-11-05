package initialize

import (
	"ams/dapi/config"
	"ams/dapi/httpserver"
	"ams/dapi/o/auth/session"
	"ams/dapi/o/org/user"
	"fmt"
	"log"
	"util/runtime"

	"github.com/jinzhu/gorm"
)

type DatabaseServer struct {
	db *gorm.DB
}

func (databaseServer *DatabaseServer) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		databaseServer.db, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbName)
			log.Fatal("Error connection:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DbName)
		}
	}

	databaseServer.db.Debug().AutoMigrate(&user.User{}, &session.Session{}) //database migration
}

func Start(p *config.ProjectConfig) {
	runtime.MaxProc()
	server = httpserver.NewProjectHttpServer(p)
}

func Wait() {
	defer beforeExit()
	server.Wait()
}
