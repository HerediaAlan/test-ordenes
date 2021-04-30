package modelos

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Nombre            string `sql:"type:NVARCHAR(30) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"nombre"`
	PrimerApellido    string `sql:"type:NVARCHAR(30) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"primerApellido"`
	SegundoApellido   string `sql:"type:NVARCHAR(30) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"segundoApellido"`
	Domicilio         string `sql:"type:NVARCHAR(100) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"domicilio"`
	Ciudad            string `sql:"type:NVARCHAR(100) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"ciudad"`
	EntidadFederativa string `sql:"type:NVARCHAR(50) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"entidadFederativa"`
	Telefono          string `sql:"type:VARCHAR(17)" json:"telefono"`
	Email             string `sql:"type:VARCHAR(50)" json:"email"`
	Ordenes           []Orden
}

func (c *Cliente) TableName() string {
	return "cliente"
}
