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
	router.GET("/students/:id", controllers.SearchStudentById)
	router.DELETE("/students/:id", controllers.DeleteStudent)
	router.PATCH("/students/:id", controllers.UpdateStudent)
	router.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	router.Run(":8080")
}
