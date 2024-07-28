package main

import (
	"fmt"
	"kalorize-api/config"
	"kalorize-api/routes"
)

func main() {
	db, err := config.InitDB()

	if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
	config.AutoMigration(db)

	// Route
	route, e := routes.Init()

	routes.RouteAuth(route, db)
	routes.RouteMakanan(route, db)
	routes.RouteQuestionnaire(route, db)
	routes.RoutesAdmin(route, db)
	routes.RouteUser(route, db)
	routes.RoutePhotoStatic(route)
	routes.RouteImportDatabase(route, db)
	routes.RouteGym(route, db)

	// Start server
	port := 80
	address := fmt.Sprintf("0.0.0.0:%d", port)
	e.Logger.Fatal(e.Start(address))
}
