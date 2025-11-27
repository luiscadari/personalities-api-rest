package main

import (
	"fmt"

	"github.com/luiscadari/personalities-api-rest/models"
	"github.com/luiscadari/personalities-api-rest/routes"
)

func main(){
	models.Personalities = []models.Personality {
		{Id: 1, Name: "First Person", History: "first history"},
	}
	fmt.Println("Hello World")
	routes.HandleRequest()
}