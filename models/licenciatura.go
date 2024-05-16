package models

import "gorm.io/gorm"

type Licenciatura struct {
	gorm.Model
	NombreLic      string         `gorm:"type:varchar(100);not null;unique_index" json:"nombre"`
	Especialidades []Especialidad `gorm:"foreignKey:LicenciaturaID" json:"especialidades"` //La clave foránea hacia LicenciaturaID en la tabla Especialidad
}
