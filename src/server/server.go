package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/justincremer/go-orm/src/book"
	"github.com/justincremer/go-orm/src/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome from http://localhost:8080")
}

func router(app *fiber.App) {
	app.Get("/", welcome)
	app.Get("/book", book.ListBooks)
	app.Get("/book/:id", book.GetBook)
	app.Post("/book", book.CreateBook)
	app.Patch("/book/:id", book.UpdateBook)
	app.Delete("/book/:id", book.DeleteBook)
}

func Create(port string, config fiber.Config) *fiber.App {
	app := fiber.New(config)

	dsn := "host=localhost user=admin password=admin dbname=admin port=5432 sslmode=disable TimeZone=America/Los_Angeles"
	conn := database.Connect(dsn)

	conn.AutoMigrate(&book.Book{})
	fmt.Println("Migration successful")

	router(app)

	return app
}
