package models

import "gorm.io/gorm"

type Reticula struct {
	gorm.Model
	EspecialidadID uint   `json:"especialidad_id"` // Clave foránea hacia Especialidad
	Semestre1      string `gorm:"not null" json:"semestre_1"`
	Semestre2      string `gorm:"not null" json:"semestre_2"`
	Semestre3      string `gorm:"not null" json:"semestre_3"`
	Semestre4      string `gorm:"not null" json:"semestre_4"`
	Semestre5      string `gorm:"not null" json:"semestre_5"`
	Semestre6      string `gorm:"not null" json:"semestre_6"`
	Semestre7      string `gorm:"not null" json:"semestre_7"`
	Semestre8      string `gorm:"not null" json:"semestre_8"`
	Semestre9      string `gorm:"not null" json:"semestre_9"`
	MateriaID      uint   `json:"materias_id"` //Clave foránea hacia Materia
}
