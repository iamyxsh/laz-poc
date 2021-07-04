package main

import (
	"LazarusPoC/configs"
	"LazarusPoC/database"
	router "LazarusPoC/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnDB()

	app := fiber.New()
	router.MainRouter(app)

	app.Listen(configs.PORT)
}
