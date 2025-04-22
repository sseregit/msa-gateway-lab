package app

import (
	"context"
	"go-gateway/app/router"
	"go.uber.org/fx"
	"log"
)

type App struct {
	router map[string]router.Router
}

func NewApp(lc fx.Lifecycle, router map[string]router.Router) App {
	a := App{router: router}

	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			go func() {
				for _, r := range router {
					if err := r.Run(); err != nil {
						panic(err)
					}
				}
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			log.Println("lifeCycle ended", c.Err())
			return nil
		},
	})

	return a
}
