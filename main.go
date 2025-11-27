package main

import (
	"fmt"

	"github.com/luiscadari/personalities-api-rest/routes"
)

func main(){
	fmt.Println("Hello World")
	routes.HandleRequest()
}