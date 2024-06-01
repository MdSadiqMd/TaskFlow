package routes

import (
	"context"
	"github.com/MdSadiqMd/TaskFlow/db"
	"github.com/MdSadiqMd/TaskFlow/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo

	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return c.Status(500).SendString("Failed to retrieve todos")
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
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

	return c.Status(201).JSON(todos)
}
