package delayed

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maodematos/hi-gofiber/pkg/util"
)

func CreateRoutes(r fiber.Router) {
	delayed := r.Group("/delay")

	delayed.Get("/:delay<int>?", func(ctx *fiber.Ctx) error {
		delay, _ := strconv.Atoi(ctx.Params("delay"))
		time.Sleep(time.Duration(delay) * time.Second)
		return ctx.JSON(
			util.ApiResponse{"response": fmt.Sprintf("You waited %d seconds.", delay)},
		)
	})
}
