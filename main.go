package main

import (
	"github.com/rafaelmgr12/gin-api-rest/database"
	"github.com/rafaelmgr12/gin-api-rest/routes"
)

func main() {
	database.DatabaseConnection()

	routes.HandleRequest()
}
