package config

// Gorm DB configure
// See https://medium.com/@_ektagarg/golang-a-todo-app-using-gin-980ebb7853c8
// and https://medium.com/canopas/golang-gorm-with-mysql-and-gin-ab876f406244

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE ordenes")
	if err != nil {
		db.Exec("CREATE DATABASE ordenes")
	}
}
