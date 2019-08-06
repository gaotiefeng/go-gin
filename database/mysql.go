package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init()  {
	var err  error
	SqlDB, err = sql.Open("mysql", "root:tf2019@tcp(192.168.41.99:3306)/gin?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
