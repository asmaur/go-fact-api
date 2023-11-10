package handlers

import (
	"time"

	"github.com/codeinceo/maur-trivia/database"
	"github.com/codeinceo/maur-trivia/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Home(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(fiber.StatusOK).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(fact)
}

func GetFact(c *fiber.Ctx) error {
	id := c.Params("id")

	var fact models.Fact

	result := database.DB.Db.First(&fact, "id = ?", id)

	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"fact": fact}})

}

func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")

	var payload *models.Fact

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var fact models.Fact
	result := database.DB.Db.First(&fact, "id = ?", id)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Answer != "" {
		updates["answer"] = payload.Answer
	}

	if payload.Question != "" {
		updates["question"] = payload.Question
	}

	updates["updated_at"] = time.Now()

	database.DB.Db.Model(&fact).Updates(updates)
	return c.Status(fiber.StatusOK).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Db.Delete(&models.Fact{}, "id = ?", id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
