package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/gin-api-rest/database"
	"github.com/rafaelmgr12/gin-api-rest/models"
)

func ShowStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Whats up " + name + "?",
	})
}

func CreateNewStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)

}
