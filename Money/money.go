package money

import (
	"github.com/gofiber/fiber"
)

// GetMoney gnn s
func GetMoney(c *fiber.Ctx) {
	c.Send("All Money")
}
