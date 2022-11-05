/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package main

import (
	"fmt"
	"log"
	"os"

	"fetch-rewards-backend/datastore"
	"fetch-rewards-backend/routes"

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
	println(datastore.Tq.Len())
	fmt.Printf("Count: %d\n", datastore.Tq.GetCount())

	// datastore.InitTQ("Lakshman", 300, "2022-11-11T10:00:00Z")
	// datastore.Tq.Insert("Test", 1200, "2022-10-31T10:00:00Z")
	// datastore.Tq.Insert("Test 1", 200, "2022-10-21T10:00:00Z")
	// println(datastore.Tq.Len())
	// fmt.Printf("Count: %d\n", datastore.Tq.GetCount())
	// datastore.Tq.PrintTQ()
	// datastore.InitTQ("Lakshman", 300, "2022-11-21T10:00:00Z")
	// datastore.Tq.Insert("Test 1", 200, "2022-10-21T11:00:00Z")
	// println(datastore.Tq.Len())
	// datastore.Tq.Pop()
	// println(datastore.Tq.GetOldestTransaction().Payer)
	// println(datastore.Tq.Len())
	// fmt.Printf("Count: %d\n", datastore.Tq.GetCount())
	// datastore.Tq.PrintTQ()
	// println(datastore.Tq.Len())
	// fmt.Printf("Count: %d\n", datastore.Tq.GetCount())
	routes.InitRouter(r)
	fmt.Printf("Running on Port %s \n", port)
	// Initialze Server on Port 3001
	r.Run(":" + port)
}
