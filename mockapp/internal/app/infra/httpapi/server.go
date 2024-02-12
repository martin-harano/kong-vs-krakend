package httpapi

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func NewServer(
	lifecycle fx.Lifecycle,
	router *fiber.App,
) *fasthttp.Server {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Info("Starting Server...")
				addr := fmt.Sprintf(":%s", "8080")
				if err := router.Listen(addr); err != nil {
					log.Fatalf("Error starting the server: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping the server...")
			return router.ShutdownWithContext(ctx)
		},
	})
	return router.Server()
}
