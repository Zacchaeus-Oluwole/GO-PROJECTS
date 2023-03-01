package main

import (
    "github.com/gofiber/fiber/v2"
    "api_go_docker/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    setupRoutes(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Dev Zacch!")
    })

    app.Listen(":3000")
}