package main

import (
	"fmt"
	"log"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func main() {

 	var db *sql.DB
 	// connection properties
 	config := mysql.Config {
 		User:	"wordup-admin",
 		Passwd:	"wordup-admin",
 		Net:	"tcp",
 		Addr:	"127.0.0.1:3306",
		DBName: "wordup_wordbank",
		AllowNativePasswords: true,
 	}

	// database handle
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("connection established!")


}
