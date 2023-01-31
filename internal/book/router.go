package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(r fiber.Router) {
	books := r.Group("/books")

	books.Post("/", func(ctx *fiber.Ctx) error {
		newBook, err := Service.CreateOne(ctx)
		if err != nil {
			return err
		}

		return ctx.JSON(newBook)
	})

	books.Get("/", func(ctx *fiber.Ctx) error {
		books, err := Service.GetAll(ctx)
		if err != nil {
			return err
		}

		ctx.Set("X-Total-Count", strconv.Itoa(len(*books)))
		return ctx.JSON(books)
	})

	books.Get("/:id", func(ctx *fiber.Ctx) error {
		bookId := ctx.Params("id")
		book, err := Service.GetOne(ctx, bookId)
		if err != nil {
			return err
		}

		return ctx.JSON(book)
	})
}
