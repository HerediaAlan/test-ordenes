package main

import (
	"test-ordenes/config"
	"test-ordenes/controladores"
	"test-ordenes/rutas"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	controladores.BuildGormDB()
	route := gin.Default()

	r := rutas.InicializarRutasClientes(route)
	r = rutas.InicializarRutasOrdenes(route)

	r.Run(":10040")
}
