package repository

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DrNikita/CarRental_GO.git/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CarsRepository struct {
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

func (r *CarsRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/cars")
	api.Post("/create", r.CreateCar)
	api.Delete("/delete/:id", r.DeleteCar)
	api.Get("/get/:id", r.GetCarById)
	api.Get("/get/all", r.GetCars)
}

func (r *CarsRepository) CreateCar(context *fiber.Ctx) error {
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

func (r CarsRepository) GetCars(context *fiber.Ctx) error {
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

func (r *CarsRepository) DeleteCar(context *fiber.Ctx) error {
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

func (r *CarsRepository) GetCarById(context *fiber.Ctx) error {
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
