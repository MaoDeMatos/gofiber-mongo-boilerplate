package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type BookService struct{}

var Service *BookService

func init() {
	Service = new(BookService)
}

func (BookService) CreateOne(ctx *fiber.Ctx) (*Book, error) {
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

	if err := mgm.Coll(&Book{}).SimpleFind(books, bson.M{}); err != nil {
		return nil, err
	}

	return books, nil
}
