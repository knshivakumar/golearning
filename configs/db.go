package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load env file variables
// connecting to databases
func ConnectDB() *sql.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Unable to load env file %v", errEnv)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	//Open connections to DB
	DB, err := sql.Open("mysql", dsn)
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error in connecting to database: %v", err)
	}
	log.Println("Successfully Connected to Database")

	return DB
}
