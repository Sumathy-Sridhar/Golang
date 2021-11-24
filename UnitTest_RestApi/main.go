package main

// commands to run test ==>
//go test
// go test -v
// go test -cover

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func PostMethod(c *gin.Context) {
	fmt.Println("\napi.go 'PostMethod' called")
	message := "PostMethod called"
	c.JSON(http.StatusOK, message)

	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetMethod(c *gin.Context) {
	fmt.Println("\napi.go 'GetMethod' called")
	message := "GetMethod called"
	c.JSON(http.StatusOK, message)
	c.IndentedJSON(http.StatusOK, albums)
}

// func PostMethod(c *gin.Context) {
// 	fmt.Println("\napi.go 'PostMethod' called")
// 	message := "PostMethod called"
// 	c.JSON(http.StatusOK, message)
// }

// func GetMethod(c *gin.Context) {
// 	fmt.Println("\napi.go 'GetMethod' called")
// 	message := "GetMethod called"
// 	c.JSON(http.StatusOK, message)
// }

func PutMethod(c *gin.Context) {
	fmt.Println("\napi.go 'PutMethod' called")
	message := "PutMethod called"
	c.JSON(http.StatusOK, message)
}

func main() {
	router := gin.Default()

	router.POST("/postdata", PostMethod)
	router.GET("/getdata", GetMethod)
	router.PUT("/putdata", PutMethod)

	listenPort := "4000"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)
}
