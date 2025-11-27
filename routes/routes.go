package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luiscadari/personalities-api-rest/controller"
)

func HandleRequest(){
	gorilla := mux.NewRouter()
	gorilla.HandleFunc("/", controller.Home)
	gorilla.HandleFunc("/personalities", controller.GetAllPersonalities).Methods("Get")
	gorilla.HandleFunc("/personalities/{id}", controller.GetPersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", gorilla))
}

