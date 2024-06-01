package db

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Collection *mongo.Collection

func ConnectDB() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	if MONGODB_URI == "" {
		log.Fatal("MONGODB_URI is not set in .env file")
		return fmt.Errorf("MONGODB_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Connected to MongoDB")
	Collection = client.Database("Tasks").Collection("todos")
	return nil
}

func ConnectDBHandler(c *fiber.Ctx) error {
	err := ConnectDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to connect to MongoDB")
	}
	return c.Status(fiber.StatusOK).SendString("Connected to MongoDB")
}
