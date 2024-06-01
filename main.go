package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	/* "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" */
	"go.mongodb.org/mongo-driver/mongo"
	/* "go.mongodb.org/mongo-driver/mongo/options" */
	"log"
	"os"
)

var collection *mongo.Collection

type Todo struct {
	ID        int    `json:"id" bson="_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World")
	err := godotenv.Load("env")
	if err != nil {
		log.Fatal("Error in loading .env File")
	}
	PORT := os.Getenv("PORT")
	/* var name string = "Md"
	const name1 string = "sadiq"
	name2 := "Mohammad"
	fmt.Println(name, name1, name2) */
	/* var x int = 5
	var p *int = &x
	fmt.Println(p)
	fmt.Println(*p) */
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})
	log.Fatal(app.Listen(":" + PORT))
}
