package router

import (
	"fmt"
	"hi-gofiber/types"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddDelayedRoutes(r fiber.Router) {
	delayed := r.Group("/delay")

	delayed.Get("/:delay<int>?", func(ctx *fiber.Ctx) error {
		delay, _ := strconv.Atoi(ctx.Params("delay"))
		time.Sleep(time.Duration(delay) * time.Second)
		return ctx.JSON(
			types.ApiResponse{"response": fmt.Sprintf("You waited %d seconds.", delay)},
		)
	})
}
