package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simonqiang/go_crash_course/book"
)

func helloWorld(c *fiber.Ctx) error {
	c.Send()
	return c.SendString("hello world")
}

func setupBooks(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.HewBooks)
	app.Delete("/api/v1/book", book.DeleteBooks)
}

func main() {
	app := fiber.New()

	setupBooks(app)
	app.Get("/", helloWorld)

	app.Listen(":3000")

}
