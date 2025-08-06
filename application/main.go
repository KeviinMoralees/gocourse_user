package main

import (
	"fmt"
	"gocourse_user/infrastructure/entrypoints/rest"
	"log"
	"net/http"
)

func main() {
	fmt.Println("🚀 Starting User Service...")

	// Setup routes
	mux := rest.SetupRoutes()

	// Start server
	port := ":8081"
	fmt.Printf("📡 Server running on port %s\n", port)
	fmt.Println("📋 Available endpoints:")
	fmt.Println("   GET /users/hello")
	fmt.Println("   GET /users/goodbye")

	log.Fatal(http.ListenAndServe(port, mux))
}
