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
	fechas            := [5]string{"Mon, 26 Apr 2021 20:12:00 MST", "Mon, 26 Apr 2021 21:33:20 MST", "Mon, 26 Apr 2021 23:45:32 MST", "Mon, 27 Apr 2021 00:01:30 MST", "Tue, 27 Apr 2021 00:02:11 MST"}
	fechasComentarios := [5]string{"Mon, 27 Apr 2021 20:14:00 MST", "Mon, 27 Apr 2021 21:35:22 MST", "Mon, 28 Apr 2021 23:47:21 MST", "Mon, 29 Apr 2021 00:03:02 MST", "Tue, 30 Apr 2021 00:04:58 MST"}
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
			t, err := time.Parse(time.RFC1123, fechas[n])
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
			tc, err := time.Parse(time.RFC1123, fechasComentarios[c])
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