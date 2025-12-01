package main

import (
	"fmt"

	"github.com/luiscadari/personalities-api-rest/database"
	"github.com/luiscadari/personalities-api-rest/models"
	"github.com/luiscadari/personalities-api-rest/routes"
)

func main(){
	models.Personalities = []models.Personality {
		{Id: 1, Name: "First Person", History: "first history"},
	}
	database.ConnectDB()
	fmt.Println("API listening on :8000")
	routes.HandleRequest()
}