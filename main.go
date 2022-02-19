package main

import (
	"golang-api-rest/database"
	"golang-api-rest/models"
	"golang-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Nome 1", History: "Historia 1"},
		{Id: 2, Name: "Nome 2", History: "Historia 2"},
	}
	database.ConnectToDatabase()
	routes.HandleRequest()
}
