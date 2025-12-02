package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luiscadari/personalities-api-rest/controller"
	"github.com/luiscadari/personalities-api-rest/middleware"
)

func HandleRequest(){
	gorilla := mux.NewRouter()
	gorilla.Use(ContentTypeMiddleware)
	gorilla.HandleFunc("/", controller.Home)
	gorilla.HandleFunc("/personalities", controller.GetAllPersonalities).Methods("Get")
	gorilla.HandleFunc("/personalities/{id}", controller.GetPersonality).Methods("Get")
	gorilla.HandleFunc("/personalities", controller.PostPersonality).Methods("Post")
	gorilla.HandleFunc("/personalities/{id}", controller.DeletePersonality).Methods("Delete")
	gorilla.HandleFunc("/personalities/{id}", controller.PutPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", gorilla))
}

