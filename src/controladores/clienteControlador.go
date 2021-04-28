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

func (gormDB *GormDB) ObtenerClientesSegunID(c *gin.Context) {
	clientes := []modelos.Cliente{}
	id := c.Param("ID")

	if err := gormDB.DB.Find(&clientes, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H {"status": http.StatusOK, "data": clientes})
}

func (gormDB *GormDB) CrearCliente(c *gin.Context) {
	var (
		cliente modelos.Cliente
		result gin.H
	)
	cliente.Nombre            = c.PostForm("nombre")
	cliente.PrimerApellido    = c.PostForm("primer_apellido")
	cliente.SegundoApellido   = c.PostForm("segundo_apellido")
	cliente.Domicilio         = c.PostForm("domicilio")
	cliente.Ciudad            = c.PostForm("ciudad")
	cliente.EntidadFederativa = c.PostForm("entidad_federativa")
	cliente.Telefono          = c.PostForm("telefono")
	cliente.Email             = c.PostForm("email")
	gormDB.DB.Create(&cliente)

	result = gin.H {
		"status": 200,
		"result": cliente,
	}
	c.JSON(http.StatusOK, result)
}

func (gormDB *GormDB) EliminarCliente(c *gin.Context) {
	var (
		cliente modelos.Cliente
		result gin.H
	)
	id := c.Param("ID")
	err := gormDB.DB.First(&cliente, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H {"status": http.StatusNotFound, "data": "Cliente no encontrado"})
	}

	err = gormDB.DB.Where("ID = ?", id).Delete(&cliente).Error
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "Error al eliminar cliente",
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "Cliente eliminado",
		}
	}

	c.JSON(http.StatusOK, result)
}
