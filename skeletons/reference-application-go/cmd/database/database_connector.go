package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

var postgres *sql.DB

func SetupDatabase() {
	connectionString := fmt.Sprintf(DatabaseConnectionString, getPostgresPassword(), getPostgresHost())

	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error occurred while opening sql connection, err: %v", err)
	}
	_, err = database.Exec(CreateDatabaseQuery)

	if err != nil {
		log.Fatalf("Error occurred while creating the counter database, error: %v", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalf("Error occurred while checking database liveness, error: %v", err)
	}

	_, err = database.Exec(CreateTableQuery)
	if err != nil {
		log.Fatalf("Error occurred while creating the counter table, error: %v", err)
	}

	log.Info("Successfully connected to the database")
	postgres = database
}

func getPostgresPassword() string {
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return "password"
	}
	return postgresPassword
}

func getPostgresHost() string {
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		return "localhost"
	}
	return postgresHost
}

func GetDatabase() *sql.DB {
	return postgres
}
