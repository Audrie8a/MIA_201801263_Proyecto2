package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

var DB *sql.DB

func Connect() {
	//var err error
	db, err := sql.Open("godror", "audrie/audrie@localhost:1521/ORCLCDB.localdomain")
	DB = db
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer db.Close()

}
