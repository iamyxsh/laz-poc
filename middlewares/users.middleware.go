package middlewares

import (
	"LazarusPoC/database"
	"LazarusPoC/middlewares/validators"
	"LazarusPoC/models"

	"github.com/gofiber/fiber/v2"
)

func CheckUniqueEmail(c *fiber.Ctx) error {
	var alreadyUser models.User
	body := c.Locals("body").(*validators.CreateUserReq)

	database.DB.Where(&models.User{Email: body.Email}).First(&alreadyUser)

	if alreadyUser.Email != "" {
		return c.JSON(fiber.Map{
			"msg": "Already a user.",
		})
	} else {
		return c.Next()
	}
}
