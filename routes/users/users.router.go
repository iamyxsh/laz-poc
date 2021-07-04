package userrouter

import (
	usercontroller "LazarusPoC/controllers/users"
	"LazarusPoC/middlewares"
	"LazarusPoC/middlewares/validators"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.Router) {
	u := *app

	u.Get("/", usercontroller.GetAllUsers)
	u.Post("/", validators.CheckCreateUserReq, middlewares.CheckUniqueEmail, usercontroller.CreateUser)
	u.Put("/:id", usercontroller.SyncUser)
}
