package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/maodematos/hi-gofiber/pkg/util"

	"go.mongodb.org/mongo-driver/bson"
)

type BookService struct{}

var Service *BookService

func init() {
	Service = new(BookService)
}

func (BookService) Create(ctx *fiber.Ctx) (*Book, error) {
	newBook := new(Book)
	coll := mgm.Coll(newBook)

	if err := ctx.BodyParser(newBook); err != nil {
		return nil, err
	}

	coll.Create(newBook)

	return newBook, nil
}

func (BookService) GetAll(ctx *fiber.Ctx) (*[]Book, error) {
	books := &[]Book{}

	queryOptions, _ := util.GetMongoOptionsFromQuery(ctx)

	if err := mgm.Coll(&Book{}).SimpleFind(books, bson.M{}, queryOptions); err != nil {
		return nil, err
	}

	return books, nil
}

func (BookService) GetOne(ctx *fiber.Ctx, id string) (*Book, error) {
	book := new(Book)

	if err := mgm.Coll(book).FindByID(id, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (BookService) PatchOne(ctx *fiber.Ctx, book *Book) (*Book, error) {
	if err := ctx.BodyParser(book); err != nil {
		return nil, err
	}

	if err := mgm.Coll(book).Update(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (BookService) DeleteOne(ctx *fiber.Ctx, book *Book) error {
	if err := mgm.Coll(book).Delete(book); err != nil {
		return err
	}

	return nil
}
