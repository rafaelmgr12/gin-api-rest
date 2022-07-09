package main

import (
	"github.com/rafaelmgr12/gin-api-rest/database"
	"github.com/rafaelmgr12/gin-api-rest/models"
	"github.com/rafaelmgr12/gin-api-rest/routes"
)

func main() {
	database.DatabaseConnection()
	models.Students = []models.Student{
		{Name: "Rafael", CPF: "123456789", RG: "123456789"},
		{Name: "Mariana", CPF: "123456789", RG: "123456789"},
	}

	routes.HandleRequest()
}
