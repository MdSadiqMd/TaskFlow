package routes

import ( /*
		"context"*/
	"github.com/MdSadiqMd/TaskFlow/db"
	"github.com/gofiber/fiber/v2"
	/* "go.mongodb.org/mongo-driver/bson"
	"log" */)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func GetTodos(c *fiber.Ctx) error {
	app := fiber.New()
	app.Get("/connectDB", db.ConnectDB)
	var todos []Todo
	cursor,err:=db.ConnectDB().collection.Find()
}
