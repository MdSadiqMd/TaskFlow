package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"Sadiq/TaskFlow/routes"
	"log"
	"os"
)


func main() {
	fmt.Println("Hello World")
	/* var name string = "Md"
	const name1 string = "sadiq"
	name2 := "Mohammad"
	fmt.Println(name, name1, name2) */
	/* var x int = 5
	var p *int = &x
	fmt.Println(p)
	fmt.Println(*p) */
	err := godotenv.Load("env")
	if err != nil {
		log.Fatal("Error in loading .env File")
	}
	PORT := os.Getenv("PORT")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})
	app.Get("/api/todos", routes.getTodos)
	log.Fatal(app.Listen("0.0.0.0" + PORT))
}
