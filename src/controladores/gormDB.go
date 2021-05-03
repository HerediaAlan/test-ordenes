package controladores

import (
	"fmt"
	"time"
	"strconv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/brianvoe/gofakeit/v6"
	"test-ordenes/modelos"
)

type GormDB struct {
	DB *gorm.DB
}

func BuildGormDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@/ordenes?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("[GORM] Failed to connect to database")
	}

	db.AutoMigrate(&modelos.Cliente{})
	db.AutoMigrate(&modelos.Orden{})
	db.AutoMigrate(&modelos.OrdenComentario{})

	// Crear datos falsos para testing
	// Se hace uso de la libraría GoFakeIt
	var (
		cliente         modelos.Cliente
		orden           modelos.Orden
		ordenComentario modelos.OrdenComentario
	)
	fechas            := [5]string{"Fri, 23 Apr 2021 20:12:00 GMT", "Sat, 24 Apr 2021 21:33:20 GMT", "Mon, 26 Apr 2021 14:45:32 GMT", "Mon, 26 Apr 2021 20:01:30 GMT", "Mon, 26 Apr 2021 21:05:11 GMT"}
	fechasComentarios := [5]string{"Sat, 24 Apr 2021 08:14:00 GMT", "Sat, 24 Apr 2021 17:35:22 GMT", "Wed, 28 Apr 021 00:47:21 GMT", "Wed, 28 Apr 2021 10:03:02 GMT", "Thu, 29 Apr 2021 00:04:58 GMT"}
	ids               := [5]string{"1", "1", "2", "2", "2"}
	asuntos           := [5]string{"Testing", "Compra testing", "Asunto de prueba", "Prueba", "Esto es una prueba"}
	descripciones     := [5]string{"Esto es una descripcion", "Descripcion", "Es descripcion", "Testing", "Descripcion de compra"}
	descripcionesSeg  := [5]string{"Comentario", "Seguimiento test", "Testing de seguimiento", "Es un seguimiento", "Prueba"}

	gofakeit.Seed(0)
	// Crear datos si no existe un cliente
	if err := db.First(&cliente).Error; err != nil {
		for i := 0; i < 15; i++ {
			cliente = modelos.Cliente{
				Nombre: gofakeit.FirstName(), PrimerApellido: gofakeit.LastName(), 
				SegundoApellido: gofakeit.LastName(), Domicilio: fmt.Sprintf("%s %s", gofakeit.StreetName(), gofakeit.StreetNumber()), 
				Ciudad: gofakeit.City(), EntidadFederativa: gofakeit.State(), 
				Telefono: gofakeit.Phone(), Email: gofakeit.Email(),
			}
			db.Create(&cliente)
		}

		for n := 0; n < 5; n++ {
			// Convertir tiempo al formato RFC1123
			// Wed, 02 Feb 2019 10:15:06 MST
			// https://apuntes.de/golang/time-parseo-de-fechas/#gsc.tab=0
			t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", fechas[n])
			if err == nil {
				idCliente, err := strconv.Atoi(ids[n])
				if err == nil {
					orden = modelos.Orden{
						FechaCreacion: t,
						ClienteID:     idCliente,
						Asunto:        asuntos[n],
						Descripcion:   descripciones[n],
					}
					db.Create(&orden)
				}
			}
		}

		for c := 0; c < 5; c++ {
			tc, err := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", fechasComentarios[c])
			if err == nil {
				idOrden, err := strconv.Atoi(ids[c])
				if err == nil {
					ordenComentario = modelos.OrdenComentario{
						FechaCreacion:          tc,
						OrdenID:                idOrden,
						DescripcionSeguimiento: descripcionesSeg[c],
					}
					db.Create(&ordenComentario)
				}
			}
		}
	}

	return db
}