package main

import (
	"github.com/ahhzaky/Golang-Gorm-And-Fiber/book"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBooks)

}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)
}
