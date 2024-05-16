package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ARDAV1D/go_apirest-tecnm/db"
	"github.com/ARDAV1D/go_apirest-tecnm/models"
	"github.com/gorilla/mux"
)

func GetLicenciaturasHandler(w http.ResponseWriter, r *http.Request) {
	var licenciaturas []models.Licenciatura
	db.DB.Find(&licenciaturas)
	json.NewEncoder(w).Encode(&licenciaturas)
}

func GetLicenciaturaHandler(w http.ResponseWriter, r *http.Request) {
	var licenciatura models.Licenciatura
	params := mux.Vars(r)
	db.DB.First(&licenciatura, params["id"])

	if licenciatura.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Licenciatura not found"))
		return
	}

	db.DB.Model(&licenciatura).Association("Especialidades").Find(&licenciatura.Especialidades)

	json.NewEncoder(w).Encode(&licenciatura)
}

func PostLicenciaturaHandler(w http.ResponseWriter, r *http.Request) {
	var licenciatura models.Licenciatura
	json.NewDecoder(r.Body).Decode(&licenciatura)

	createdlicenciatura := db.DB.Create(&licenciatura)
	err := createdlicenciatura.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&licenciatura)
}

func PutLicenciaturaHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la licenciatura de la URL
	params := mux.Vars(r)
	licenciaturaID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID de licenciatura no válido"))
		return
	}

	// Decodificar el cuerpo de la solicitud en una estructura Licenciatura
	var licenciatura models.Licenciatura
	if err := json.NewDecoder(r.Body).Decode(&licenciatura); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Actualizar la licenciatura en la base de datos
	licenciatura.ID = uint(licenciaturaID) // Asignar el ID de la licenciatura
	updatedLicenciatura := db.DB.Save(&licenciatura)
	if err := updatedLicenciatura.Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Devolver la licenciatura actualizada como respuesta
	json.NewEncoder(w).Encode(&licenciatura)
}

func DeleteLicenciaturaHandler(w http.ResponseWriter, r *http.Request) {
	var licenciatura models.Licenciatura
	params := mux.Vars(r)
	db.DB.First(&licenciatura, params["id"])

	if licenciatura.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Licenciatura not found"))
		return
	}

	// Eliminar la licenciatura de la base de datos
	db.DB.Unscoped().Delete(&licenciatura)

	// Reiniciar la secuencia de ID
	if err := db.DB.Exec("ALTER SEQUENCE licenciaturas_id_seq RESTART WITH 1").Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
