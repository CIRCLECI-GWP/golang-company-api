package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Tech Company listing API with Golang"})
}

func main() {
	router := gin.Default()
	router.GET("/", HomepageHandler)
	router.Run()
}

type Company struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	CEO     string `json:"ceo"`
	Revenue string `json:"revenue"`
}
