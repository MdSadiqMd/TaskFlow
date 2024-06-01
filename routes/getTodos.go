package routes

import (
	"context"
	"github.com/MdSadiqMd/TaskFlow/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return c.Status(500).SendString("Failed to retrieve todos")
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Fatal(err)
			return c.Status(500).SendString("Failed to decode todo")
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return c.Status(500).SendString("Cursor error")
	}

	return c.JSON(todos)
}
