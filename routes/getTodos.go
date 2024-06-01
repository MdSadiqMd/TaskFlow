package routes

import (/* 
	"context"
	"github.com/MdSadiqMd/TaskFlow/db" */
	"github.com/gofiber/fiber/v2"
	/* "go.mongodb.org/mongo-driver/bson"
	"log" */
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func GetTodos(c *fiber.Ctx) error { 
	return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
}
