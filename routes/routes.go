package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/gin-api-rest/controllers"
)

func HandleRequest() {
	router := gin.Default()
	router.GET("/students", controllers.ShowStudents)
	router.GET("/hello/:name", controllers.Hello)
	router.POST("/students", controllers.CreateNewStudent)
	router.Run(":8080")
}
