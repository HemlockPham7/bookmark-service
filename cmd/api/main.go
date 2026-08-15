package main

import "github.com/HemlockPham7/bookmark-service/internal/infrastructure"

// @title Bookmark Service API
// @version 1.0.0
// @description This is the API documentation for the Bookmark service.
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Init api
	app := infrastructure.CreateAPI()

	// Run api
	err := app.Start()
	if err != nil {
		panic(err)
	}
}
