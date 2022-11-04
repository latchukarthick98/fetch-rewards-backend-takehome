package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {

	// Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	// Initialize Gin
	r := gin.Default()

	port := getEnvVar("PORT")
	fmt.Printf("Port %s \n", port)
	// Initialze Server on Port 3001
	r.Run(":" + port)
}
