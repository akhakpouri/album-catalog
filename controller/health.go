package controller

import (
	"mock-api/managers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	healths := managers.Get()
	c.IndentedJSON(http.StatusOK, healths)
}
