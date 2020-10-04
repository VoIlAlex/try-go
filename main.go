package main

import (
	"try-go/api/view"
)

func main() {
	// Set mode to release
	//gin.SetMode(gin.ReleaseMode)
	view.StartServer()
}
