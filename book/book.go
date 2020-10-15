package book

import (
	"github.com/gofiber/fiber"
)

func GetBook(c *fiber.Ctx) {
	c.Send("A single Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adds a new book")
}

func DeleteBooks(c *fiber.Ctx) {
	c.Send("Delete Books")
}
