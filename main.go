package main

import (
	"net/http"

	"github.com/ARDAV1D/go_apirest-tecnm/db"
	"github.com/ARDAV1D/go_apirest-tecnm/models"
	"github.com/ARDAV1D/go_apirest-tecnm/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Licenciatura{})
	db.DB.AutoMigrate(models.Especialidad{})
	db.DB.AutoMigrate(models.Materia{})
	db.DB.AutoMigrate(models.Reticula{})

	r := mux.NewRouter()

	//Licenciatura routes
	r.HandleFunc("/licenciaturas", routes.GetLicenciaturasHandler).Methods("GET")
	r.HandleFunc("/licenciaturas", routes.PostLicenciaturaHandler).Methods("POST")
	r.HandleFunc("/licenciaturas/{id}", routes.GetLicenciaturaHandler).Methods("GET")
	r.HandleFunc("/licenciaturas/{id}", routes.PutLicenciaturaHandler).Methods("PUT")
	r.HandleFunc("/licenciaturas/{id}", routes.DeleteLicenciaturaHandler).Methods("DELETE")

	//Especialidad routes
	r.HandleFunc("/especialidads", routes.GetEspecialidadesHandler).Methods("GET")
	r.HandleFunc("/especialidads", routes.PostEspecialidadHandler).Methods("POST")
	r.HandleFunc("/especialidads/{id}", routes.GetEspecialidadHandler).Methods("GET")
	r.HandleFunc("/especialidads/{id}", routes.PutEspecialidadHandler).Methods("PUT")
	r.HandleFunc("/especialidads/{id}", routes.DeleteEspecialidadHandler).Methods("DELETE")

	//Materia routes
	r.HandleFunc("/materias", routes.GetMateriasHandler).Methods("GET")
	r.HandleFunc("/materias", routes.PostMateriaHandler).Methods("POST")
	r.HandleFunc("/materias/{id}", routes.GetMateriaHandler).Methods("GET")
	r.HandleFunc("/materias/{id}", routes.PutMateriaHandler).Methods("PUT")
	r.HandleFunc("/materias/{id}", routes.DeleteMateriaHandler).Methods("DELETE")

	//Reticula routes
	r.HandleFunc("/reticulas", routes.GetReticulasHandler).Methods("GET")
	r.HandleFunc("/reticulas", routes.PostReticulaHandler).Methods("POST")
	r.HandleFunc("/reticulas/{id}", routes.GetReticulaHandler).Methods("GET")
	r.HandleFunc("/reticulas/{id}", routes.PutReticulaHandler).Methods("PUT")
	r.HandleFunc("/reticulas/{id}", routes.DeleteReticulaHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
