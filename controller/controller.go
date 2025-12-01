package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luiscadari/personalities-api-rest/database"
	"github.com/luiscadari/personalities-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request){
	var p []models.Personality
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonality(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var p []models.Personality
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}