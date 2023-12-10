package middlaware

import (
	"go-nat-project/models"
	"go-nat-project/utils"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin(c *fiber.Ctx) error {

	token := c.Get("Authorization")

	if token != "1bd51b70-f679-4396-a416-67607010b192" {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4001,
			ErrorData: models.ApiError{
				ErrorTitle:   "Invalid Token",
				ErrorMessage: "Invalid Token",
			},
		})
	}

	return c.Next()
}
