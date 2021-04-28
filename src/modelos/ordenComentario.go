package modelos

import (
	"time"
	"gorm.io/gorm"
)

type OrdenComentario struct {
	gorm.Model
	FechaCreacion          time.Time `sql:type:"DATETIME NOT NULL CHARACTER SET utf8 COLLATE utf8_general_ci" json:"fechaCreacion"`
	OrdenID                int
	DescripcionSeguimiento string `sql:type:"NVARCHAR(100) NOT NULL CHARACTER SET utf8 COLLATE utf8_general_ci" json:"descripcionSeguimiento"`
}

func (c *OrdenComentario) TableName() string {
	return "orden_comentario"
}
