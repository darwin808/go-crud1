package book

import (
	"test1/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)

}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)

	if book.Title == "" {
		return c.Status(500).JSON(&fiber.Map{

			"success": false,
			"error":   "no books found",
		})
	}
	return c.JSON(&fiber.Map{"success": true, "data": book})
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.SendStatus(503)
	}

	db.Create(&book)

	return c.SendString("Object created")
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book

	db.First(&book, id)
	if book.Title == "" {
		return c.SendStatus(500)
	}

	db.Delete(&book)
	return c.SendString("Book successfully deleted")
}
