package routes

import (
	"log"
	"net/http"

	"github.com/luiscadari/personalities-api-rest/controller"
)

func HandleRequest(){
	http.HandleFunc("/", controller.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}