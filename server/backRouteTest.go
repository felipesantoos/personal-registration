package server

import (
	"log"
	"net/http"
	"personal-registration/routers"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func OpenServerTest() {
	router := mux.NewRouter()

	// Configuração do CORS
	//c := cors.AllowAll()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001", "http://localhost:3000"}, // Origens permitidas (adapte ao seu ambiente)
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	//PERSON ROUTERS
	// Defina uma rota para lidar com solicitações POST para criar um paciente
	router.HandleFunc("/api/person", routers.CreatePerson).Methods(http.MethodPost)

	router.HandleFunc("/api/person/{id}", routers.EditPerson).Methods(http.MethodPut)

	router.HandleFunc("/api/person/{id}", routers.GetPersonById).Methods(http.MethodGet)

	router.HandleFunc("/api/personDelete/{id}", routers.DeletePerson).Methods(http.MethodGet)
	//
	//router.HandleFunc("/api/patients/{id}", routers.EditPerson).Methods(http.MethodPut)
	//
	//router.HandleFunc("/api/patients/login", routers.ValidateLoginPatient).Methods(http.MethodPost)
	//
	//router.HandleFunc("/api/patients/deactivate", routers.DeletePerson).Methods(http.MethodPost)

	// server static files
	fs := http.FileServer(http.Dir("./tmp"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	handler := c.Handler(router)

	println("Servidor ligado na porta :8000!")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
