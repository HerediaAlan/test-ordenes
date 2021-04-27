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
		db.Exec("USE ordenes")
/*
		query := `CREATE TABLE IF NOT EXISTS cliente(
				clienteID INT AUTO_INCREMENT PRIMARY KEY,
				nombre VARCHAR(30) NOT NULL,
				primerApellido VARCHAR(30) NOT NULL,
				segundoApellido VARCHAR(30) NOT NULL,
				domicilio VARCHAR(100),
				ciudad VARCHAR(100),
				entidadFederativa VARCHAR(50),
				telefono VARCHAR(17),
				email VARCHAR(50)
			)`
		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}

		queryInsert := `INSERT INTO cliente(nombre, primerApellido, segundoApellido, 
							domicilio, ciudad, entidadFederativa, telefono, email) VALUES
							("José", "Pérez", "Sánchez", "Av. Jalisco y 26", "San Luis Río Colorado", 
								"Sonora", "6530005030", "jose@gmail.com"),
							("Adriana", "García", "Ibarra", "Av. Tlaxcala 23 y 24", "San Luis Río Colorado",
								"Sonora", "6539998877","adriana.garcia@gmail.com")`

		_, err1 := db.Exec(queryInsert)
		if err1 != nil {
			panic(err1)
		}*/
	}
}
