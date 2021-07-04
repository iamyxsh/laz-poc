package usercontroller

import (
	"LazarusPoC/database"
	"LazarusPoC/helpers"
	"LazarusPoC/middlewares/validators"
	"LazarusPoC/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	body := c.Locals("body").(*validators.CreateUserReq)
	user := models.User{Name: body.Name, Email: body.Email, Mfa: body.MFADetails.MFA, Mfa_type: body.MFADetails.MFAType, Phone: body.Phone, Country: body.Country}
	result := database.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.JSON(fiber.Map{
		"msg": "User Created Succesfully.",
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User

	err := database.DB.Find(&users)
	if err != nil {
		return c.JSON(fiber.Map{
			"msg": err,
		})
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func SyncUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var awUser models.AWUser

	res, err := helpers.Resty.R().Get(fmt.Sprintf("/user/%v", id))
	if err != nil {
		return c.JSON(fiber.Map{
			"msg": err,
		})
	}
	json.Unmarshal(res.Body(), &awUser)

	result := database.DB.Model(&models.User{}).Where("email", awUser.Email).Updates(models.User{Name: awUser.Name, Email: awUser.Email})
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"msg": err,
		})
	}

	return c.JSON(fiber.Map{
		"msg": "User Udated Succesfully.",
	})
}
