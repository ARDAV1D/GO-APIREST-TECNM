package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ARDAV1D/go_apirest-tecnm/db"
	"github.com/ARDAV1D/go_apirest-tecnm/models"
	"github.com/gorilla/mux"
)

func GetReticulasHandler(w http.ResponseWriter, r *http.Request) {
	var reticulas []models.Reticula
	db.DB.Find(&reticulas)
	json.NewEncoder(w).Encode(&reticulas)
}

func GetReticulaHandler(w http.ResponseWriter, r *http.Request) {
	var reticula models.Reticula
	params := mux.Vars(r)
	db.DB.First(&reticula, params["id"])

	if reticula.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Reticula not found"))
		return
	}
	json.NewEncoder(w).Encode(&reticula)
}
func PostReticulaHandler(w http.ResponseWriter, r *http.Request) {
	var Reticula models.Reticula
	json.NewDecoder(r.Body).Decode(&Reticula)

	createdreticula := db.DB.Create(&Reticula)
	err := createdreticula.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Reticula)
}

func PutReticulaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reticulaID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID de reticula no válido"))
		return
	}

	var reticula models.Reticula
	if err := json.NewDecoder(r.Body).Decode(&reticula); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	reticula.ID = uint(reticulaID)

	if err := db.DB.Save(&reticula).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&reticula)
}

func DeleteReticulaHandler(w http.ResponseWriter, r *http.Request) {
	var reticula models.Reticula
	params := mux.Vars(r)
	db.DB.First(&reticula, params["id"])

	if reticula.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Reticula not found"))
		return
	}

	db.DB.Unscoped().Delete(&reticula)

	if err := db.DB.Exec("ALTER SEQUENCE reticulas_id_seq RESTART WITH 1").Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
