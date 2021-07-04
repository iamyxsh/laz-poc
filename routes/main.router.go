package router

import (
	orgrouter "LazarusPoC/routes/organizations"
	userrouter "LazarusPoC/routes/users"

	"github.com/gofiber/fiber/v2"
)

func MainRouter(app *fiber.App) {
	api := app.Group("api")

	users := api.Group("users")
	userrouter.Router(&users)

	orgs := api.Group("orgs")
	orgrouter.Router(&orgs)
}
