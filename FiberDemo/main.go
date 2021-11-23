package main

import (
	book "FiberDemo/Book"

	"github.com/gofiber/fiber"
)

func fibdemo(c *fiber.Ctx) {
	c.Send("Fiber Framework Demo!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", fibdemo)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(8080)
}

// func main() {
// 	app := fiber.New()
// 	app.Get("/", func(c *fiber.Ctx) {
// 		c.Send("Fiber Framework Demo")
// 	})
// 	app.Listen(8080)
// }
