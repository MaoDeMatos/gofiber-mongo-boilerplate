package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(r fiber.Router) {
	books := r.Group("/books")

	books.Post("/", func(ctx *fiber.Ctx) error {
		newBook, err := Service.Create(ctx)
		if err != nil {
			return err
		}

		return ctx.Status(fiber.StatusCreated).JSON(newBook)
	})

	books.Get("/", func(ctx *fiber.Ctx) error {
		books, err := Service.GetAll(ctx)
		if err != nil {
			return err
		}

		// TODO: Get the actual total count instead of this
		ctx.Set("X-Total-Count", strconv.Itoa(len(*books)))
		return ctx.JSON(books)
	})

	booksById := books.Group("/:id<string>", func(ctx *fiber.Ctx) error {
		bookId := ctx.Params("id")
		book, err := Service.GetOne(ctx, bookId)

		if err != nil {
			return err
		}

		ctx.Locals("currentBook", book)
		return ctx.Next()
	})

	booksById.Get("/", func(ctx *fiber.Ctx) error {
		book := ctx.Locals("currentBook")

		return ctx.JSON(book)
	})

	booksById.Patch("/", func(ctx *fiber.Ctx) error {
		book := ctx.Locals("currentBook")
		book, err := Service.PatchOne(ctx, book.(*Book))
		if err != nil {
			return err
		}

		return ctx.JSON(book)
	})

	booksById.Delete("/", func(ctx *fiber.Ctx) error {
		book := ctx.Locals("currentBook")
		err := Service.DeleteOne(ctx, book.(*Book))
		if err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	})
}
