package book

import (
	"github.com/gofiber/fiber"
)

// GetBook gnn s
func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

// GetBook gnn
func GetBook(c *fiber.Ctx) {
	c.Send("A single Book")
}

// NewBook gnn
func NewBook(c *fiber.Ctx) {
	c.Send("Adds a new book")
}

// DeleteBooks gnn
func DeleteBooks(c *fiber.Ctx) {
	c.Send("Delete Books")
}
