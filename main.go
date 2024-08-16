package main

import (
	"api/album-catalog/internal/album"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("album", addAlbum)
	router.Run("localhost:8080")
}

func addAlbum(c *gin.Context) {
	var bufferAlbum album.Album
	err := c.BindJSON(&bufferAlbum)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	album.AddAlbum(bufferAlbum)
	c.IndentedJSON(http.StatusCreated, bufferAlbum)
}

func getAlbums(c *gin.Context) {
	albums := album.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("id"))
	if convErr != nil {
		panic(convErr)
	}
	album, err := album.GetAlbumById(id)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		c.Abort()
	} else {
		c.IndentedJSON(http.StatusOK, album)
	}
}
