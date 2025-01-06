// main.go
package main

import (
    "log"
    "net/http"
    "Assignment_One/handlers"
    "Assignment_One/middleware"
    "Assignment_One/db"
)

func main() {
    // Initialize database
    database, err := db.InitDB("blog.db")
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    // Initialize router
    router := http.NewServeMux()
    
    // Initialize handler with database
    h := handlers.NewHandler(database)

    // Routes
    router.HandleFunc("/blog", middleware.Chain(h.CreateBlog, 
        middleware.LogRequest,
        middleware.ValidateJSON,
        middleware.Authenticate))
    
    router.HandleFunc("/blog/", middleware.Chain(h.HandleBlog,
        middleware.LogRequest,
        middleware.Authenticate))
        
    router.HandleFunc("/blogs", middleware.Chain(h.ListBlogs,
        middleware.LogRequest))

    log.Println("Server starting on :8082")
    log.Fatal(http.ListenAndServe(":8082", router))
}