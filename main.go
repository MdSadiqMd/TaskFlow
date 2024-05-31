package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World")
	/* var name string = "Md"
	const name1 string = "sadiq"
	name2 := "Mohammad"
	fmt.Println(name, name1, name2) */
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})
	log.Fatal(app.Listen(":3000"))
}
