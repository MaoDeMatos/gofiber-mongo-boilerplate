package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(r fiber.Router) {
	books := r.Group("/books")

	books.Post("/", createBook)
	books.Get("/", getAllBooks)
}

func createBook(ctx *fiber.Ctx) error {
	newBook, err := Service.CreateOne(ctx)
	if err != nil {
		return err
	}
	return ctx.JSON(newBook)
}

func getAllBooks(ctx *fiber.Ctx) error {
	books, err := Service.GetAll(ctx)
	if err != nil {
		return err
	}

	ctx.Set("X-total-count", strconv.Itoa(len(*books)))
	return ctx.JSON(books)
}
