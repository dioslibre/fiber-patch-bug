package main

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func test(c *fiber.Ctx) {
	fmt.Println(c.Method())
	c.Status(200).JSON(c.Method())
}

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Post("/api", test)

	app.Patch("/api", test)

	app.Listen(3000)
}
