package routes

import (
	"golang-api-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
