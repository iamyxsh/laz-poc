package orgrouter

import (
	orgcontroller "LazarusPoC/controllers/organizations"
	"LazarusPoC/middlewares/validators"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.Router) {
	u := *app

	u.Get("/:id", orgcontroller.GetAllOrgByUserId)
	u.Post("/:id", validators.CheckCreateOrgReq, orgcontroller.CreateOrg)
	u.Put("/:id", validators.CheckCreateOrgReq, orgcontroller.UpdateOrgById)
}
