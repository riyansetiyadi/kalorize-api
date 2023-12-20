package main

import (
	"fmt"
	"kalorize-api/config"
	"kalorize-api/routes"
)

func main() {
	db := config.InitDB()
	config.AutoMigration(db)

	// Route
	route, e := routes.Init()
	routes.RouteAuth(route, db)

	// Start server
	port := 8080
	address := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(address))
}
