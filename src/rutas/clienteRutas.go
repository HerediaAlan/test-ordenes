package rutas

// https://github.com/ilham-dev/go-gin-mvc-crud/blob/master/route/users.go

import (
	"test-ordenes/controladores"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InicializarRutas() *gin.Engine {
	db := controladores.BuildGormDB()
	gormDB := &controladores.GormDB{DB: db}
	router := gin.Default();

	router.Use(cors.Default())

	router.GET("/clientes", gormDB.ObtenerClientes)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H {"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return router;
}
