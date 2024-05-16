package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ARDAV1D/go_apirest-tecnm/db"
	"github.com/ARDAV1D/go_apirest-tecnm/models"
	"github.com/gorilla/mux"
)

func GetEspecialidadesHandler(w http.ResponseWriter, r *http.Request) {
	var especialidads []models.Especialidad
	db.DB.Find(&especialidads)
	json.NewEncoder(w).Encode(&especialidads)
}

func GetEspecialidadHandler(w http.ResponseWriter, r *http.Request) {
	var especialidad models.Especialidad
	params := mux.Vars(r)
	db.DB.First(&especialidad, params["id"])

	if especialidad.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Especialidad not found"))
		return
	}
	db.DB.Model(&especialidad).Association("Materias").Find(&especialidad.Materias)

	json.NewEncoder(w).Encode(&especialidad)
}
func PostEspecialidadHandler(w http.ResponseWriter, r *http.Request) {
	var Especialidad models.Especialidad
	json.NewDecoder(r.Body).Decode(&Especialidad)

	createdespecialidad := db.DB.Create(&Especialidad)
	err := createdespecialidad.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Especialidad)
}

func PutEspecialidadHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la especialidad de la URL
	params := mux.Vars(r)
	especialidadID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID de especialidad no válido"))
		return
	}

	// Decodificar el cuerpo de la solicitud en una estructura Especialidad
	var especialidad models.Especialidad
	if err := json.NewDecoder(r.Body).Decode(&especialidad); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Asignar el ID de la especialidad
	especialidad.ID = uint(especialidadID)

	// Actualizar la especialidad en la base de datos
	if err := db.DB.Save(&especialidad).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Devolver la especialidad actualizada como respuesta
	json.NewEncoder(w).Encode(&especialidad)
}

func DeleteEspecialidadHandler(w http.ResponseWriter, r *http.Request) {
	var especialidad models.Especialidad
	params := mux.Vars(r)
	db.DB.First(&especialidad, params["id"])

	if especialidad.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Especialidad not found"))
		return
	}

	// Eliminar la especialidad de la base de datos
	db.DB.Unscoped().Delete(&especialidad)

	// Reiniciar la secuencia de ID para Especialidad
	if err := db.DB.Exec("ALTER SEQUENCE especialidads_id_seq RESTART WITH 1").Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
