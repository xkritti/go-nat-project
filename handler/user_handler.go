package user_handler

import (
	"go-nat-project/database"
	user_models "go-nat-project/models"

	"github.com/gofiber/fiber/v2"
)

// can get all user and filter by year and level_type and major_type
func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []user_models.Users
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No user found!",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"data":    users,
			"message": "Get all user success!",
		})
	}
}
