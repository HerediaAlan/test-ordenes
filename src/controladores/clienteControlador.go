package controladores

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
