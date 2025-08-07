package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KeviinMoralees/gocourse_user/infrastructure/entrypoints/rest/handlers/read"
	"github.com/KeviinMoralees/gocourse_user/infrastructure/entrypoints/rest/handlers/writer"
	"github.com/gorilla/mux"
)

func main() {
	// Get role from environment variable, default to "all"
	role := os.Getenv("APP_ROLE")
	if role == "" {
		role = "all"
	}

	fmt.Println("🚀 Starting User Service...")
	fmt.Printf("🎭 Role: %s\n", role)

	// Validate role at startup
	validateRole(role)

	// Setup routes based on role
	router := setupRoutes(role)

	// Start server
	port := ":8081"
	fmt.Printf("📡 Server running on port %s\n", port)
	fmt.Println("📋 Available endpoints:")

	switch role {
	case "writer":
		fmt.Println("   POST /users")
		fmt.Println("   PUT  /users/{id}")
		fmt.Println("   DELETE /users/{id}")
	case "read":
		fmt.Println("   GET  /users")
		fmt.Println("   GET  /users/{id}")
	case "all":
		fmt.Println("   GET  /users")
		fmt.Println("   GET  /users/{id}")
		fmt.Println("   POST /users")
		fmt.Println("   PUT  /users/{id}")
		fmt.Println("   DELETE /users/{id}")
	}

	log.Fatal(http.ListenAndServe(port, router))
}

// setupRoutes configures all REST API routes using gorilla/mux based on role
func setupRoutes(role string) *mux.Router {
	router := mux.NewRouter()

	// Register routes based on role
	switch role {
	case "read":
		fmt.Println("📖 Setting up READ-only routes...")
		read.SetupReadRoutes(router)
	case "writer":
		fmt.Println("✍️ Setting up WRITE-only routes...")
		writer.SetupWriterRoutes(router)
	case "all":
		fmt.Println("🔄 Setting up ALL routes...")
		read.SetupReadRoutes(router)
		writer.SetupWriterRoutes(router)
	default:
		fmt.Println("⚠️ Invalid role, defaulting to ALL routes...")
		read.SetupReadRoutes(router)
		writer.SetupWriterRoutes(router)
	}

	return router
}

// validateRole validates the role and provides feedback
func validateRole(role string) {
	validRoles := []string{"read", "writer", "all"}

	isValid := false
	for _, validRole := range validRoles {
		if role == validRole {
			isValid = true
			break
		}
	}

	if !isValid {
		fmt.Printf("⚠️ Warning: Invalid role '%s'. Valid roles are: %v\n", role, validRoles)
		fmt.Println("🔄 Defaulting to 'all' role...")
	}
}
