package routes

import (	
	"Assignment_two/controllers"
	"Assignment_two/middleware"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	
	router.Use(middleware.Authenticate)
	router.Use(middleware.ValidateJSON)

	
	router.HandleFunc("/product", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
	

	return router
}
