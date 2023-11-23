package main

import (
	"fmt"
	"log"
	"test1/book"
	"test1/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {

	return c.SendString("Hello, World 👋!")
}

func helloDarwin(c *fiber.Ctx) error {
	return c.Status(200).SendString("ok")

}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/darwin", helloDarwin)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connectin opened to database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated!!!")
}

func main() {
	app := fiber.New()
	initDatabase()

	app.Use(cors.New())
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
