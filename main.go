package main

import (
	"log"
	"os"

	"github.com/DrNikita/CarRental_GO.git/models"
	"github.com/DrNikita/CarRental_GO.git/repository"
	"github.com/DrNikita/CarRental_GO.git/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// DB connectiong....
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.Connect(config)

	if err != nil {
		log.Fatal("Could not load the database")
	}
	// connected successfully)

	err = models.MigrateCars(db)
	if err != nil {
		log.Fatal(err)
	}
	err = models.MigrateRoles(db)
	if err != nil {
		log.Fatal(err)
	}
	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	carsRepo := repository.CarsRepository{DB: db}

	app := fiber.New()
	carsRepo.SetupRoutes(app)

	app.Listen(":8080")
}
