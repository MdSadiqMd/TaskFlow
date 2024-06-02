package main

import (
	"fmt"
	"github.com/MdSadiqMd/TaskFlow/db"
	"github.com/MdSadiqMd/TaskFlow/routes"
	"github.com/gofiber/fiber/v2"
	/* "github.com/gofiber/fiber/v2/middleware/cors" */
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
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3000"
	} else {
		PORT = ":" + PORT
	}

	app := fiber.New()

	/* app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	})) */

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	app.Get("/api/todos", routes.GetTodos)
	app.Post("/api/todos", routes.CreateTodo)
	app.Patch("/api/todos/:id", routes.UpdateTodo)
	app.Delete("/api/todos/:id", routes.DeleteTodo)

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0"+PORT))
}
