package orgcontroller

import (
	"LazarusPoC/database"
	"LazarusPoC/middlewares/validators"
	"LazarusPoC/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateOrg(c *fiber.Ctx) error {
	userid, _ := strconv.ParseInt(c.Params("id"), 10, 32)
	body := c.Locals("body").(*validators.CreateOrgReq)
	org := models.Organization{
		Name:     body.Name,
		Type:     body.Type,
		Address:  body.Address,
		City:     body.City,
		Country:  body.Country,
		Timezone: body.Timezone,
		Status:   body.Status,
	}

	result := database.DB.Create(&org)

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	fmt.Println(userid)

	member := models.Membership{
		UserID:         int(userid),
		OrganizationID: int(org.ID),
		Name:           org.Name,
		Role:           "owner",
	}

	result = database.DB.Create(&member)

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.JSON(fiber.Map{
		"msg": "Org created successfully.",
	})
}

func UpdateOrgById(c *fiber.Ctx) error {
	userid, _ := strconv.ParseInt(c.Params("id"), 10, 32)
	body := c.Locals("body").(*validators.CreateOrgReq)
	org := models.Organization{
		Name:     body.Name,
		Type:     body.Type,
		Address:  body.Address,
		City:     body.City,
		Country:  body.Country,
		Timezone: body.Timezone,
		Status:   body.Status,
	}

	result := database.DB.Model(&models.Organization{}).Where("id", userid).Updates(&org)

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	return c.JSON(fiber.Map{
		"msg": "Org updated successfully.",
	})
}

func GetAllOrgByUserId(c *fiber.Ctx) error {
	userid, _ := strconv.ParseInt(c.Params("id"), 10, 32)

	var memberships []models.Membership
	var orgs []models.Organization

	result := database.DB.Where("user_id", userid).Find(&memberships)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"msg": result.Error,
		})
	}

	for _, member := range memberships {
		var o models.Organization

		database.DB.Where("id", member.OrganizationID).First(&o)
		orgs = append(orgs, o)
	}

	return c.JSON(fiber.Map{
		"orgs": orgs,
	})
}
