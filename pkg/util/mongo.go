package util

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoOptionsFromQuery(ctx *fiber.Ctx) (*options.FindOptions, error) {
	queryOptions := new(options.FindOptions)
	if err := ctx.QueryParser(queryOptions); err != nil {
		return nil, err
	}

	sortBy := ctx.Query("sort", "")
	var sortDirection int
	if strings.HasPrefix(sortBy, "-") {
		sortBy = sortBy[1:]
		sortDirection = -1
	} else {
		sortDirection = 1
	}

	queryOptions.SetSort(bson.D{{Key: sortBy, Value: sortDirection}})

	return queryOptions, nil
}
