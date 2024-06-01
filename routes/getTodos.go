package routes

import ("github.com/gofiber/fiber/v2"
"db\db.go")

type Todo struct {
	ID        int    `json:"id" bson="_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func getTodos(c *fiber.Ctx) error{
	var todos []Todo
}