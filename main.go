package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DrNikita/CarRental_GO.git/models"
	"github.com/DrNikita/CarRental_GO.git/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Car struct {
	Id         uint      `json:"id"`
	Govnum     string    `json:"govnum"`
	Brand      string    `json:"brand"`
	IssueDate  time.Time `json:"issue_date"`
	CarCost    uint      `json:"car_cost"`
	RentalCost uint      `json:"rental_cost"`
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_car", r.CreateCar)
	api.Delete("/delete_car/:id", r.DeleteCar)
	api.Get("/get_car/:id", r.GetCarById)
	api.Get("/cars", r.GetCars)
}

func (r *Repository) CreateCar(context *fiber.Ctx) error {
	car := Car{}
	err := context.BodyParser(&car)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&car).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create car"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "car has been created"})
	return nil
}

func (r *Repository) GetCars(context *fiber.Ctx) error {
	carModels := &[]models.Cars{}

	err := r.DB.Find(carModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the books"})
		return err
	}
	context.Status(fiber.StatusOK).JSON(
		&fiber.Map{
			"message": "car fetched successfully",
			"data":    carModels,
		})
	return nil
}

func (r *Repository) DeleteCar(context *fiber.Ctx) error {
	carModel := models.Cars{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := r.DB.Delete(carModel, id)

	if err.Error != nil {
		context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete car",
		})
		return err.Error
	}
	context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "car delete successfully",
	})
	return nil
}

func (r *Repository) GetCarById(context *fiber.Ctx) error {
	carModel := &models.Cars{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is ", id)

	err := r.DB.Where("id = ?", id).First(carModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the car"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "car id fetched successfully",
		"data":    carModel,
	})
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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

	err = models.MigrateCars(db)
	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
}
