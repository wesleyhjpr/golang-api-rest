package main

import (
	"golang-api-rest/models"
	"golang-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Nome 1", History: "Historia 1"},
		{Name: "Nome 2", History: "Historia 2"},
	}
	routes.HandleRequest()
}
