package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB() // connect to database

	routes.UserRoute(app) // define user routes

	app.Listen(":6000")
}
