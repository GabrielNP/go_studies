package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gin_and_tests/controllers"
	"gin_and_tests/database"
	"gin_and_tests/models"
)

var ID int

func setupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // minimalist test response
	routes := gin.Default()
	return routes
}

func createStudentMock() {
	student := models.Student{
		Name:                "Name Test 1",
		FederalRegistration: "12345678901",
		DocumentNumber:      "12345678X",
	}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func deleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student.ID)
}

func TestCheckStatusCodeGreetRequestWithParam(t *testing.T) {
	database.Connect()
	createStudentMock()
	// defer deleteStudentMock()
	r := setupTestRoutes()
	r.GET("/:name", controllers.Greet)
	req, _ := http.NewRequest("GET", "/gabriel", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	// if resp.Code != http.StatusOK {
	// 	t.Fatalf("Status error: expecting %d, got %d", http.StatusOK, resp.Code)
	// }

	assert.Equal(t, http.StatusOK, resp.Code, "should be equals")

	responseMock := `{"API says:":"Hey gabriel, sup?"}`
	responseBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, responseMock, string(responseBody))
}
