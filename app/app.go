package app

import (
	"album-catalog/controller"

	"github.com/gin-gonic/gin"
)

type App struct {
	Routes *gin.Engine
}

func (a *App) CreateRoutes() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumById)
	router.POST("/album", controller.AddAlbum)
	router.Run()
	a.Routes = router
}

func (a *App) Run() {
	a.Routes.Run(":8080")
}
