package api

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/Yanuarprayoga9/GO-BLOG-RESTFUL-API/api/controllers"
    "github.com/Yanuarprayoga9/GO-BLOG-RESTFUL-API/api/seed"
    "github.com/joho/godotenv"
    "github.com/gorilla/handlers"
)

var server = controllers.Server{}

func Run() {

    var err error
    err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error getting env, not comming through %v", err)
    } else {
        fmt.Println("We are getting the env values")
    }

    server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

    seed.Load(server.DB)

    // Define CORS options
    corsOptions := handlers.AllowedOrigins([]string{"*"})
    corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
    corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

    // Wrap your router with CORS middleware
    handler := handlers.CORS(corsOptions, corsMethods, corsHeaders)(server.Router)

    log.Fatal(http.ListenAndServe(":8080", handler))
}
