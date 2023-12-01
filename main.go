package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type albumn struct {
	ID     string  `json : "id"`
	Title  string  `json : "title"`
	Artist string  `json : "artist"`
	Price  float64 `json : "float"`
}

var albumns = []albumn{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbumns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumns)
}

func getAlbumnsById(c *gin.Context) {
	id := c.Param("id")
	for _, albumn := range albumns {
		if albumn.ID == id {
			c.IndentedJSON(http.StatusOK, albumn)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albumn not found"})
}

func postAlbumn(c *gin.Context) {
	var newAlbumn albumn
	if err := c.BindJSON(&newAlbumn); err != nil {
		return
	}
	albumns = append(albumns, newAlbumn)
	c.IndentedJSON(http.StatusCreated, newAlbumn)
}

func main() {
	router := gin.Default()
	router.GET("/albumns", getAlbumns)
	router.GET("/albumns/:id", getAlbumnsById)
	router.POST("/albumns", postAlbumn)

	router.Run("localhost:8080")
}
