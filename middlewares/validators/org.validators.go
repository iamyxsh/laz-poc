package validators

import (
	"LazarusPoC/helpers"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateOrgReq struct {
	Name     string    `json:"name" validate:"required"`
	Type     string    `json:"type" validate:"required"`
	Address  string    `json:"address" validate:"required"`
	City     string    `json:"city"`
	Country  string    `json:"country"`
	Timezone time.Time `json:"timezone"`
	Status   string    `json:"status"`
}

func CheckCreateOrgReq(c *fiber.Ctx) error {
	b := new(CreateOrgReq)
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
