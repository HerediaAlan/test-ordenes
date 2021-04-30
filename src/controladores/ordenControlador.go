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

	// https://gorm.io/docs/preload.html
	if err := gormDB.DB.Preload("OrdenComentarios").Find(&ordenes).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H {"status": http.StatusOK, "data": ordenes})
}

func (gormDB *GormDB) ObtenerOrdenSegunID(c *gin.Context) {
	orden := modelos.Orden{}
	id := c.Param("ID")

	if err := gormDB.DB.First(&orden, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H {"status": http.StatusOK, "data": orden})
}

func (gormDB *GormDB) CrearOrden(c *gin.Context) {
	var (
		orden  modelos.Orden
		result gin.H
	)

	// https://www.pauladamsmith.com/blog/2011/05/go_time.html
	// https://apuntes.de/golang/time-parseo-de-fechas/#gsc.tab=0
	t, err := time.Parse(time.RFC1123, c.PostForm("fecha_creacion"))
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

func (gormDB *GormDB) ActualizarOrden(c *gin.Context) {
	id := c.Query("ID")
	var (
		orden      modelos.Orden
		result     gin.H
	)

	err := gormDB.DB.Where("id = ?", id).First(&orden).Error
	if err != nil {
		result = gin.H {
			"status": 400,
			"result": "Orden especificada no encontrada",
		}
	}

	_, errF := time.Parse(time.RFC1123, c.PostForm("fecha_creacion"))
	if errF != nil {
		result = gin.H {
			"status": 400,
			"result": "Error al convertir fecha",
		}
		c.JSON(400, result)
		return
	}
	_, errId := strconv.Atoi(c.PostForm("clienteID"))
	if errId != nil {
		result = gin.H {
			"status": 400,
			"result": "Error al convertir ID de cliente",
		}
		c.JSON(400, result)
		return
	}
	c.BindJSON(&orden)

	err = gormDB.DB.Save(&orden).Error
	if err != nil {
		result = gin.H {
			"status": 400,
			"result": "No se pudo actualizar la orden",
		}
	} else {
		result = gin.H {
			"status": 200,
			"result": "Datos actualizados correctamente",
		}
	}

	c.JSON(http.StatusOK, result)
}


func (gormDB *GormDB) EliminarOrden(c *gin.Context) {
	var (
		orden  modelos.Orden
		result gin.H
	)
	id := c.Param("ID")
	err := gormDB.DB.First(&orden, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H {"status": http.StatusNotFound, "data": "Orden no encontrado"})
	}

	err = gormDB.DB.Where("id = ?", id).Delete(&orden).Error
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "Error al eliminar orden",
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "Orden eliminado",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (gormDB *GormDB) ObtenerOrdenComentarios(c *gin.Context) {
	var (
		orden  modelos.Orden
		result gin.H
	)
	comentarios := []modelos.OrdenComentario{}
	ordenID := c.Param("ID")
	err := gormDB.DB.First(&orden, ordenID).Error
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "No se encontró la orden",
		}
	} else {
		err = gormDB.DB.Where("orden_id = ?", ordenID).Find(&comentarios).Error
		if err != nil {
			result = gin.H{
				"status": 400,
				"result": "Error al buscar comentarios",
			}
		} else {
			result = gin.H {
				"status": 200,
				"result": comentarios,
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

func (gormDB *GormDB) CrearOrdenComentario(c *gin.Context) {
	var (
		orden      modelos.Orden
		comentario modelos.OrdenComentario
		result     gin.H
	)
	ordenID := c.Param("ID")
	err := gormDB.DB.First(&orden, ordenID).Error
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "No se encontró la orden",
		}
	} else {
		t, err := time.Parse(time.RFC1123, c.PostForm("fecha_creacion"))
		if err != nil {
			result = gin.H {
				"status": 400,
				"result": "Error al convertir fecha",
			}
			c.JSON(400, result)
			return
		}

		id, err := strconv.Atoi(c.PostForm("ordenID"))
		if err != nil {
			result = gin.H {
				"status": 400,
				"result": "Error al convertir ID de orden",
			}
			c.JSON(400, result)
			return
		}

		comentario.FechaCreacion          = t
		comentario.OrdenID                = id
		comentario.DescripcionSeguimiento = c.PostForm("descripcion_seguimiento")
		gormDB.DB.Create(&comentario)
		result = gin.H {
			"status": 200,
			"result": comentario,
		}
		c.JSON(http.StatusOK, result)
	}
}