package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/gin-api-rest/controllers"
	"github.com/rafaelmgr12/gin-api-rest/database"
	"github.com/rafaelmgr12/gin-api-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
}

func MockStudent() {
	student := models.Student{Name: "Student Test", CPF: "123456789", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.First(&student)
	database.DB.Delete(&student, ID)
}

func TestVerifiedHelloParamater(t *testing.T) {
	router := SetupTestRouter()
	router.GET("/hello/:name", controllers.Hello)
	req, _ := http.NewRequest("GET", "/hello/rafael", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
	answerMock := `{"API says:":"Whats up rafael?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, answerMock, string(responseBody))
}

func TestListAllStudentsHandler(t *testing.T) {
	database.DatabaseConnection()
	MockStudent()
	defer DeleteStudentMock()
	router := SetupTestRouter()
	router.GET("/students", controllers.ShowStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
}

func TestSearchStudentByCPFHandler(t *testing.T) {
	database.DatabaseConnection()
	MockStudent()
	defer DeleteStudentMock()
	router := SetupTestRouter()
	router.GET("/students/cpf/:cpf", controllers.SearchStudentByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/123456789", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be equal")
}

func TestSearchStudentByIdHandler(t *testing.T) {
	database.DatabaseConnection()
	MockStudent()
	defer DeleteStudentMock()
	router := SetupTestRouter()
	router.GET("/students/:id", controllers.SearchStudentById)
	searchPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", searchPath, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	var studentsMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentsMock)
	assert.Equal(t, "Student Test", studentsMock.Name, "Should be equal")
	assert.Equal(t, "123456789", studentsMock.CPF)
	assert.Equal(t, "123456789", studentsMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)

}

func TestDeleteStudentByHandler(t *testing.T) {
	database.DatabaseConnection()
	MockStudent()
	router := SetupTestRouter()
	router.DELETE("/students/:id", controllers.DeleteStudent)
	searchPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", searchPath, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateStudentHandler(t *testing.T) {
	database.DatabaseConnection()
	MockStudent()
	defer DeleteStudentMock()
	router := SetupTestRouter()
	router.PATCH("/students/:id", controllers.UpdateStudent)

	student := models.Student{Name: "Student Test", CPF: "473456789", RG: "123456789"}
	jsonStudent, _ := json.Marshal(student)
	searchPath := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", searchPath, bytes.NewBuffer(jsonStudent))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	var studentMockUpdate models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMockUpdate)
	assert.Equal(t, "Student Test", studentMockUpdate.Name)
	assert.Equal(t, "473456789", studentMockUpdate.CPF)
	assert.Equal(t, "123456789", studentMockUpdate.RG)

}
