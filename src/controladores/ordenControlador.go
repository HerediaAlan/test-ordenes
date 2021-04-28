package controladores

// https://github.com/ilham-dev/go-gin-mvc-crud/blob/master/controller/UserController.go

import (
	"time"
	"test-ordenes/modelos"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)


func (gormDB *GormDB) ObtenerOrdenes(c *gin.Context) {
	ordenes := []modelos.Orden{}

	if err := gormDB.DB.Find(&ordenes).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H {"status": http.StatusOK, "data": ordenes})
}

func (gormDB *GormDB) CrearOrden(c *gin.Context) {
	var (
		orden modelos.Orden
		result gin.H
	)

	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, c.PostForm("fecha_creacion"))
	if err != nil {
		result = gin.H {
			"status": 400,
			"result": "Error al convertir fecha",
		}
		c.JSON(400, result)
		return
	}
	id, err := strconv.Atoi(c.PostForm("clienteID"))
	if err != nil {
		result = gin.H {
			"status": 400,
			"result": "Error al convertir ID de cliente",
		}
		c.JSON(400, result)
		return
	}

	orden.FechaCreacion = t
	orden.ClienteID     = id
	orden.Asunto        = c.PostForm("asunto")
	orden.Descripcion   = c.PostForm("descripcion")
	gormDB.DB.Create(&orden)

	result = gin.H {
		"status": 200,
		"result": orden,
	}
	c.JSON(http.StatusOK, result)
}

/*
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

func (gormDB *GormDB) ActualizarCliente(c *gin.Context) {
	id := c.Query("id")
	var (
		cliente      modelos.Cliente
		nuevoCliente modelos.Cliente
		result       gin.H
	)

	err := gormDB.DB.First(&cliente, id).Error
	if err != nil {
		result = gin.H {
			"status": 400,
			"result": "Cliente especificado no encontrado",
		}
	}

	nuevoCliente.Nombre            = c.PostForm("nombre")
	nuevoCliente.PrimerApellido    = c.PostForm("primer_apellido")
	nuevoCliente.SegundoApellido   = c.PostForm("segundo_apellido")
	nuevoCliente.Domicilio         = c.PostForm("domicilio")
	nuevoCliente.Ciudad            = c.PostForm("ciudad")
	nuevoCliente.EntidadFederativa = c.PostForm("entidad_federativa")
	nuevoCliente.Telefono          = c.PostForm("telefono")
	nuevoCliente.Email             = c.PostForm("email")

	err = gormDB.DB.Model(&cliente).Updates(nuevoCliente).Error
	if err != nil {
		result = gin.H {
			"status": 400,
			"result": "No se pudo actualizar el cliente",
		}
	} else {
		result = gin.H {
			"status": 200,
			"result": "Datos actualizados correctamente",
		}
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

	err = gormDB.DB.Where("id = ?", id).Delete(&cliente).Error
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
*/