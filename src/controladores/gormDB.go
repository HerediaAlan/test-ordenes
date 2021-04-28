package controladores

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test-ordenes/modelos"
	"github.com/brianvoe/gofakeit/v6"
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
	var cliente modelos.Cliente

	gofakeit.Seed(0)
	if err := db.First(&cliente).Error; err != nil {
		for i := 1; i < 10; i++ {
			cliente = modelos.Cliente{
				Nombre: gofakeit.FirstName(), PrimerApellido: gofakeit.LastName(), 
				SegundoApellido: gofakeit.LastName(), Domicilio: fmt.Sprintf("%s %s", gofakeit.StreetName(), gofakeit.StreetNumber()), 
				Ciudad: gofakeit.City(), EntidadFederativa: gofakeit.State(), 
				Telefono: gofakeit.Phone(), Email: gofakeit.Email(),
			}
			db.Create(&cliente)
		}
	}

	return db
}