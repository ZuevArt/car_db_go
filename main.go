package main

import (
    "log"
    "net/http"

	"car_db/internal/utils"
    "car_db/internal/handlers"
    "car_db/internal/middleware"
)

func init() {
    utils.StartAdminKeyRotator()
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("POST /login", handlers.Login)
	mux.HandleFunc("POST /register", handlers.Register)

    api := http.NewServeMux()
    api.HandleFunc("GET /cars", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetCars)))
	api.HandleFunc("POST /cars", middleware.AdminMiddleware(handlers.CreateCar))
	api.HandleFunc("DELETE /users/{id}", middleware.AdminMiddleware(handlers.DeleteUser))

	mux.Handle("/api/", http.StripPrefix("/api", api))

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}