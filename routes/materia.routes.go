package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ARDAV1D/go_apirest-tecnm/db"
	"github.com/ARDAV1D/go_apirest-tecnm/models"
	"github.com/gorilla/mux"
)

func GetMateriasHandler(w http.ResponseWriter, r *http.Request) {
	var materias []models.Materia
	db.DB.Find(&materias)
	json.NewEncoder(w).Encode(&materias)
}

func GetMateriaHandler(w http.ResponseWriter, r *http.Request) {
	var materia models.Materia
	params := mux.Vars(r)
	db.DB.First(&materia, params["id"])

	if materia.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Materia not found"))
		return
	}

	db.DB.Model(&materia).Association("Reticulas").Find(&materia.Reticulas)

	json.NewEncoder(w).Encode(&materia)
}
func PostMateriaHandler(w http.ResponseWriter, r *http.Request) {
	var Materia models.Materia
	json.NewDecoder(r.Body).Decode(&Materia)

	createdmateria := db.DB.Create(&Materia)
	err := createdmateria.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Materia)
}

func PutMateriaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	materiaID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID de Materia no válido"))
		return
	}

	var materia models.Materia
	if err := json.NewDecoder(r.Body).Decode(&materia); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	materia.ID = uint(materiaID)

	if err := db.DB.Save(&materia).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&materia)
}

func DeleteMateriaHandler(w http.ResponseWriter, r *http.Request) {
	var materia models.Materia
	params := mux.Vars(r)
	db.DB.First(&materia, params["id"])

	if materia.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Materias not found"))
		return
	}

	db.DB.Unscoped().Delete(&materia)

	if err := db.DB.Exec("ALTER SEQUENCE materia_id_seq RESTART WITH 1").Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
