package routes

import (
	"context"
	"github.com/MdSadiqMd/TaskFlow/db"
	"github.com/MdSadiqMd/TaskFlow/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		log.Println("Error parsing body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if todo.Body == "" {
		log.Println("Empty Todo Body")
		return c.Status(400).JSON(fiber.Map{"error": "Empty Todo cannot be created"})
	}
	insertResult, err := db.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Println("Failed to insert todo:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to insert todo", "details": err.Error()})
	}
	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		todo.ID = oid
	} else {
		log.Println("Failed to cast InsertedID to ObjectID")
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get inserted ID"})
	}

	return c.Status(201).JSON(todo)
}
