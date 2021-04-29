package rutas

// https://github.com/ilham-dev/go-gin-mvc-crud/blob/master/route/users.go

import (
	"test-ordenes/controladores"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InicializarRutasOrdenes(router *gin.Engine) *gin.Engine {
	db := controladores.BuildGormDB()
	gormDB := &controladores.GormDB{DB: db}

	// Para permitir llamadas desde una webapp como Axios.js
	router.Use(cors.Default())

	router.GET("/ordenes", gormDB.ObtenerOrdenes)
	router.GET("/ordenes/:ID", gormDB.ObtenerOrdenSegunID)
	router.POST("/ordenes", gormDB.CrearOrden)
	router.PUT("/ordenes/:ID", gormDB.ActualizarOrden)
	router.DELETE("/ordenes/:ID", gormDB.EliminarOrden)

	router.GET("/ordenes/:ID/comentarios", gormDB.ObtenerOrdenComentarios)
	router.POST("/ordenes/:ID/comentarios", gormDB.CrearOrdenComentario)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H {"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return router;
}
