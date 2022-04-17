package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomepageHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	r := SetUpRouter()
	r.GET("/", HomepageHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewCompanyHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/company", NewCompanyHandler)
	companyId := xid.New().String()
	company := Company{
		ID:      companyId,
		Name:    "Demo Company",
		CEO:     "Demo CEO",
		Revenue: "35 million",
	}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/company", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
