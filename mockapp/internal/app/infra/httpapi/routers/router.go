package routers

import "github.com/gofiber/fiber/v2"

type Router interface {
	Load()
}

func MakeRouter(
	apiRouter *ApiRouter,
) *fiber.App {
	cfg := fiber.Config{
		AppName:       "API Gateway Benchmark",
		CaseSensitive: true,
		Prefork:       true,
	}

	r := fiber.New(cfg)
	apiRouter.Load(r)
	return r
}
