package main

import (
	"log"
	"net/http"

	"car_db/internal/handlers"
	"car_db/internal/middleware"
	"car_db/internal/utils"
)

func init() {
	utils.StartAdminKeyRotator()
}

func main() {
	mux := http.NewServeMux()

	// === Авторизация ===
	mux.HandleFunc("POST /login", handlers.Login)
	mux.HandleFunc("POST /register", handlers.Register)

	// === API ===
	api := http.NewServeMux()
	api.HandleFunc("POST /login", handlers.Login)
	api.HandleFunc("POST /register", handlers.Register)
	api.HandleFunc("GET /cars", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetCars)))
	api.HandleFunc("POST /cars", middleware.AdminMiddleware(handlers.CreateCar))
	api.HandleFunc("GET /models", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetModels)))
	api.HandleFunc("GET /cars/{id}/images", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetCarImages)))
	api.HandleFunc("GET /bodytypes", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetBodyTypes)))
	api.HandleFunc("GET /drivetypes", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetDriveTypes)))
	api.HandleFunc("GET /engines", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetEngines)))
	api.HandleFunc("GET /transmissions", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetTransmissions)))
	api.HandleFunc("DELETE /users/{id}", middleware.AdminMiddleware(handlers.DeleteUser))
	api.HandleFunc("DELETE /cars/{id}", middleware.AdminMiddleware(http.HandlerFunc(handlers.DeleteCar)))
	api.HandleFunc("POST /cars/{id}", middleware.AdminMiddleware(handlers.UpdateCar))
	api.HandleFunc("GET /users/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetCurrentUser)))
	api.HandleFunc("POST /users/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateProfile)))

	mux.Handle("/api/", http.StripPrefix("/api", api))

	fs := http.FileServer(http.Dir("../frontend"))
	mux.Handle("/", fs)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
