package rutas

// https://github.com/ilham-dev/go-gin-mvc-crud/blob/master/route/users.go

import (
	"test-ordenes/controladores"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InicializarRutasClientes(router *gin.Engine) *gin.Engine {
	db := controladores.BuildGormDB()
	gormDB := &controladores.GormDB{DB: db}

	// Para permitir llamadas desde una webapp como Axios.js
	router.Use(cors.Default())

	router.GET("/clientes", gormDB.ObtenerClientes)
	router.GET("/clientes/:ID", gormDB.ObtenerClienteSegunID)
	router.POST("/clientes", gormDB.CrearCliente)
	router.PUT("/clientes/:ID", gormDB.ActualizarCliente)
	router.DELETE("/clientes/:ID", gormDB.EliminarCliente)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H {"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return router;
}
