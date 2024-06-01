package main

import (
	"fmt"
	"github.com/MdSadiqMd/TaskFlow/db"
	"github.com/MdSadiqMd/TaskFlow/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3000"
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})
	err = db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	app.Get("/api/todos", routes.GetTodos)
	app.Get("/api/todos", routes.CreateTodo)
	app.Get("/api/todos", routes.UpdateTodo)
	log.Fatal(app.Listen(PORT))
}
