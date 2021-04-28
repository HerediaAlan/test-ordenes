package controladores

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	var cliente modelos.Cliente
	if err := db.First(&cliente).Error; err != nil {
		cliente1 := modelos.Cliente{Nombre: "José", PrimerApellido: "Pérez", SegundoApellido: "Sánchez", Domicilio: "Av. Jalisco y 26", Ciudad: "San Luis Río Colorado", EntidadFederativa: "Sonora", Telefono: "6530005030", Email: "jose@gmail.com"}
		cliente2 := modelos.Cliente{Nombre: "Adriana", PrimerApellido: "García", SegundoApellido: "Ibarra", Domicilio: "Av. Tlaxcala 23 y 24", Ciudad: "San Luis Río Colorado", EntidadFederativa: "Sonora", Telefono: "6539998877", Email: "adriana.garcia@gmail.com"}
		db.Create(&cliente1)
		db.Create(&cliente2)
	}

	return db
}