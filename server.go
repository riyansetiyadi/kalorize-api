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
	routes.RouteMakanan(route, db)
	routes.RouteQuestionnaire(route, db)
	routes.RoutesAdmin(route, db)
	routes.RouteUser(route, db)
	routes.RoutePhotoStatic(route)

	// Start server
	port := 8080
	address := fmt.Sprintf("0.0.0.0:%d", port)
	e.Logger.Fatal(e.Start(address))
}
