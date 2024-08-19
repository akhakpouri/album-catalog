package main

import "album-catalog/app"

func main() {
	var a app.App
	a.CreateRoutes()
	a.Run()
}
