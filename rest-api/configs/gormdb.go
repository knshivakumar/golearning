package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load env file variables
// connecting to databases
func ConnectGormDB() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Unable to load env file %v", errEnv)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	//Open connections to DB
}
