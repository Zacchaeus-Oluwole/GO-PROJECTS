package main

import (
	"api_go_docker/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App){
	app.Get("/", handlers.ListFact)

	app.Post("/fact", handlers.CreateFact)
}