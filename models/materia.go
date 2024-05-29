package models

import "gorm.io/gorm"

type Materia struct {
	gorm.Model
	NombreMa       string     `gorm:"not null" json:"nombre_Ma"`
	EspecialidadID uint       `json:"especialidad_id"`                       // CLave foránea hacia Especialidad
	Reticulas      []Reticula `gorm:"foreignKey:MateriaID" json:"reticulas"` //La clave foránea hacia MateriaID en la tabla Reticula
}
