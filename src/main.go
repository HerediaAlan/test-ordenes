package main

import (
	"test-ordenes/config"
	"test-ordenes/controladores"
	"test-ordenes/rutas"
)

func main() {
	config.InitDB()
	controladores.BuildGormDB()

	r := rutas.InicializarRutas()

	r.Run(":10040")
}
