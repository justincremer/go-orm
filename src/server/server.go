package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/justincremer/go-orm/src/database"
	"github.com/justincremer/go-orm/src/models/book"
	"github.com/justincremer/go-orm/src/models/user"
)

func Create(port string, config fiber.Config) *fiber.App {
	app := fiber.New(config)
	dsn := "host=localhost user=admin password=admin dbname=admin port=5432 sslmode=disable TimeZone=America/Los_Angeles"
	conn := database.Connect(dsn)

	conn.AutoMigrate(&book.Book{})
	conn.AutoMigrate(&user.User{})
	fmt.Println("Migration successful")

	app.Get("/", welcome)
	bookRouter(app)
	userRouter(app)

	return app
}
