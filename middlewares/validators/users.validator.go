package validators

import (
	"LazarusPoC/helpers"

	"github.com/gofiber/fiber/v2"
)

type MFADetails struct {
	MFA     bool   `json:"mfa" validate:"required"`
	MFAType string `json:"mfa_type" validate:"max=10,min=2,required"`
}

type CreateUserReq struct {
	ID         string     `json:"$id" validate:"max=20,min=3,required"`
	Name       string     `json:"name" validate:"max=20,min=3,required"`
	Email      string     `json:"email" validate:"max=20,min=3,required,email"`
	Country    string     `json:"country" validate:"max=20,min=3"`
	Phone      string     `json:"phone_num" validate:"required"`
	MFADetails MFADetails `json:"mfa_details"`
}

func CheckCreateUserReq(c *fiber.Ctx) error {
	b := new(CreateUserReq)
	helpers.CheckBodyParser(c, b)

	errors := helpers.ValidateStruct(*b)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	c.Locals("body", b)

	return c.Next()
}
