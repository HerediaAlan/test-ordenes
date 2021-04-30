package modelos

import (
	"time"
	"gorm.io/gorm"
)

type Orden struct {
	gorm.Model
	FechaCreacion      time.Time `sql:type:"DATETIME NOT NULL CHARACTER SET utf8 COLLATE utf8_general_ci" json:"fechaCreacion"`
	ClienteID          int       `gorm:"constraint:OnDelete:CASCADE"`
	Asunto             string    `sql:type:"NVARCHAR(100) NOT NULL CHARACTER SET utf8 COLLATE utf8_general_ci" json:"asunto"`
	Descripcion        string    `sql:type:"NVARCHAR(100) NOT NULL CHARACTER SET utf8 COLLATE utf8_general_ci" json:"descripcion"`
	OrdenComentarios   []OrdenComentario
}

func (c *Orden) TableName() string {
	return "orden"
}
