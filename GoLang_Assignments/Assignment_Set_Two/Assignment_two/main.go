package main

import (
	"log"
	"net/http"
	"Assignment_two/routes"
	"Assignment_two/db"
)

func main() {
	
	db := config.ConnectDatabase()
	defer db.Close()

	
	router := routes.InitRoutes()

	
	log.Println("Server is running on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", router))
}
