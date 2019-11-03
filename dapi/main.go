package main

import (
	"fmt"
	// "ams/dapi/config"
	// "ams/dapi/initialize"
	// "context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// initialize.Start(ctx, config.ReadConfig())
	// initialize.Wait()

	db, err := sql.Open("mysql", "root:idfcau1992@tcp(127.0.0.1:3306)/ams")
	fmt.Println(db)

	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT users SET email=?,password=?")
	checkErr(err)

	res, err := stmt.Exec("vuhongthaihy@gmail.com", "123456")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
