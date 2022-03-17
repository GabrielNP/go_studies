package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gin_and_tests/controllers"
)

func setupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestCheckStatusCodeGreetRequestWithParam(t *testing.T) {
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
