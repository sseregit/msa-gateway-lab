package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-gateway/app/client"
	"go-gateway/config"
	"go-gateway/types/http"
)

type Router struct {
	port string
	cfg  config.App

	engin *fiber.App

	client *client.HttpClient
}

func NewRouter(cfg config.App, clients map[string]*client.HttpClient) Router {
	r := Router{
		cfg:    cfg,
		port:   fmt.Sprintf(":%s", cfg.App.Port),
		client: clients[cfg.App.Name],
	}

	r.engin = fiber.New()
	r.engin.Use(recover.New())

	/*	r.engin.Use(cors.New(cors.Config{
			//AllowOrigins:
			//AllowMethods:
			//MaxAge:
		}))
	*/
	for _, v := range cfg.Http.Router {
		r.registerRouter(v)
	}

	return r
}

func (r Router) registerRouter(v config.Router) {
	switch v.Method {
	case http.GET:
		handler := AddGet(v, r.client)
		r.engin.Get(v.Path, handler)
	case http.POST:
		handler := AddPost(v, r.client)
		r.engin.Post(v.Path, handler)
	case http.DELETE:
		handler := AddDelete(v, r.client)
		r.engin.Delete(v.Path, handler)
	case http.PUT:
		handler := AddPut(v, r.client)
		r.engin.Put(v.Path, handler)
	default:
		panic("Failed to find router method")
	}
}

func (r Router) Run() error {
	return r.engin.Listen(r.port)
}
