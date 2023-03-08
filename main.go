package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "To pimp A butterfly", Artist: "Kendrick Lamar", Price: 60.99},
	{ID: "2", Title: "The Eminem Show", Artist: "Slim Shady", Price: 80.99},
	{ID: "3", Title: "Ready to Die", Artist: "The Notorious Big", Price: 100},
}

func main() {
	// Initialize a gin router using Default
	router := gin.Default()
	// Use the GET function to associate the GET HTTP method and /albums path with a handler function
	router.GET("/albums", getItems)
	router.POST("/albums", postItem)
	router.GET("/albums/:id", getItemByID)

	// Attach the router to http server
	router.Run("localhost:8090")
}

// getItems function that takes a gin.Context parameter.
// gin.Context is the most important part of Gin; it carries the request details, validates and serializes Json and more...
// The function first argument is the HTTP status code you want to send to the client. Here we are passing statusOK from net/http package
func getItems(c *gin.Context) {
	// call context.IndentedJSON to serialize the struct into JSON and add it to the response
	c.IndentedJSON(http.StatusOK, albums)
	// More compact JSON
	//c.JSON(http.StatusOK, albums)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums = looking for an album whose ID value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func postItem(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
