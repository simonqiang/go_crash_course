package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("A Single Book")
}

func HewBooks(c *fiber.Ctx) error {
	return c.SendString("Adds a new book")
}

func DeleteBooks(c *fiber.Ctx) error {
	return c.SendString("Delete a Books")
}
