package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m, err := migrate.New("file://migrations", os.Getenv("DB_CONNECTION_STRING"))

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil {
		log.Fatal(err)
	}

}