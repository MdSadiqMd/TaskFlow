package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello World")
	/* var name string = "Md"
	const name1 string = "sadiq"
	name2 := "Mohammad"
	fmt.Println(name, name1, name2) */
	app := fiber.New()
	log.Fatal(app.Listen(":3000"))
}
