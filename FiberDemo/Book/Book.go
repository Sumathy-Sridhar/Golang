package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("Retrives all books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("Retrieves Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adding a New Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Deleting Book")
}
