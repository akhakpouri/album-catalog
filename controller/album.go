package controller

import (
	"album-catalog/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddAlbum(c *gin.Context) {
	var bufferAlbum model.Album
	err := c.BindJSON(&bufferAlbum)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	model.AddAlbum(bufferAlbum)
	c.IndentedJSON(http.StatusCreated, bufferAlbum)
}

func GetAlbums(c *gin.Context) {
	albums := model.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
	id, convErr := strconv.Atoi(c.Param("id"))
	if convErr != nil {
		panic(convErr)
	}
	album, err := model.GetAlbumById(id)

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		c.Abort()
	} else {
		c.IndentedJSON(http.StatusOK, album)
	}
}
