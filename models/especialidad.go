package models

import "gorm.io/gorm"

type Especialidad struct {
	gorm.Model
	NombreEsp      string    `gorm:"type:varchar(100);not null;unique_index" json:"nombre_Esp"`
	LicenciaturaID uint      `json:"licenciatura_id"`                           //La clave foránea hacia Licenciatura
	Materias       []Materia `gorm:"foreignKey:EspecialidadID" json:"materias"` //La clave foránea hacia EspecialidadID en la tabla Materia
}
