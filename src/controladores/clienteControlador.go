package controladores

// https://github.com/ilham-dev/go-gin-mvc-crud/blob/master/controller/UserController.go

import (
	"test-ordenes/modelos"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (gormDB *GormDB) ObtenerClientes(c *gin.Context) {
	clientes := []modelos.Cliente{}

	if err := gormDB.DB.Find(&clientes).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H {"status": http.StatusOK, "data": clientes})
}
