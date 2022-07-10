package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafaelmgr12/gin-api-rest/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	return router
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
