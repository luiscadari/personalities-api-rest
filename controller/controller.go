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

func PostPersonality(w http.ResponseWriter, r *http.Request){
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request){
	body := mux.Vars(r)
	id := body["id"]
	var personalityToDrop models.Personality
	database.DB.Delete(&personalityToDrop, id)
	json.NewEncoder(w).Encode(personalityToDrop)
}

func PutPersonality(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var personality models.Personality
	id := params["id"]
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}