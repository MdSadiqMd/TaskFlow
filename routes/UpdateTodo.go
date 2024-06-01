package routes

import (
	"context"
	"github.com/MdSadiqMd/TaskFlow/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo ID"})
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
