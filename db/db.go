package db

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	/* "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" */
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection

func ConnectDB(c *fiber.Ctx) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading .env file")
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	if MONGODB_URI == "" {
		log.Fatal("MONGODB_URI is not set in .env file")
		return c.Status(fiber.StatusInternalServerError).SendString("MONGODB_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to connect to MongoDB")
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to ping MongoDB")
	}

	fmt.Println("Connected to MongoDB")
	collection = client.Database("Tasks").Collection("todos")
	return c.Status(fiber.StatusOK).SendString("Connected to MongoDB")
}
