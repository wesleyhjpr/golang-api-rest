package routes

import (
	"golang-api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
