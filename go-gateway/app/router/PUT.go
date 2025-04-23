package router

import (
	"github.com/gofiber/fiber/v2"
	"go-gateway/app/client"
	"go-gateway/config"
)

type put struct {
	cfg    config.Router
	client *client.HttpClient
}

func AddPut(
	cfg config.Router,
	client *client.HttpClient,
) func(c *fiber.Ctx) error {
	r := put{cfg: cfg, client: client}

	return r.handleRequest
}

func (r put) handleRequest(c *fiber.Ctx) error {

	apiResult, err := r.client.PUT(r.cfg.Path, r.cfg, c.Request().Body())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(apiResult)
}
