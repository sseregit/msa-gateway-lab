package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go-gateway/app/client"
	"go-gateway/config"
	"go-gateway/types/http"
	"strings"
)

type get struct {
	cfg    config.Router
	client client.HttpClient
}

func AddGet(
	cfg config.Router,
	client client.HttpClient,
) func(c *fiber.Ctx) error {
	r := get{cfg: cfg, client: client}

	return r.handleRequest
}

func (r get) handleRequest(c *fiber.Ctx) error {
	switch r.cfg.GetType {
	case http.QUERY:
		return r.queryType(c)
	case http.URL:
		return r.urlType(c)
	default:
		panic("Failed to find get type")
	}
}

func (r get) queryType(c *fiber.Ctx) error {
	var builder strings.Builder

	for i, k := range r.cfg.Variable {
		v := utils.CopyString(c.Query(k))

		if i == 0 {
			builder.WriteString(fmt.Sprintf("?%s=%s", k, v))
		} else {
			builder.WriteString(fmt.Sprintf("&%s=%s", k, v))
		}
	}

	fullUrl := builder.String()

	apiReesut, err := r.client.GET(fullUrl, r.cfg)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(apiReesut)
}

func (r get) urlType(c *fiber.Ctx) error {
	var builder strings.Builder
	builder.WriteString(string(c.Request().URI().Path()))

	fullUrl := builder.String()

	apiReesut, err := r.client.GET(fullUrl, r.cfg)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(apiReesut)
}
