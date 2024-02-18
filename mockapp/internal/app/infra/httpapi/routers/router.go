package routers

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/infra/config"
)

type Router interface {
	Load()
}

func MakeRouter(
	config *config.Config,
	apiRouter *ApiRouter,
) *fiber.App {
	cfg := fiber.Config{
		AppName:       "API Gateway Benchmark",
		CaseSensitive: true,
		Prefork:       config.Server.Prefork,
	}
	if config.Server.UseSonic {
		log.Info("Loading Sonic JSON into the router")
		cfg.JSONEncoder = sonic.Marshal
		cfg.JSONDecoder = sonic.Unmarshal
	}
	r := fiber.New(cfg)
	apiRouter.Load(r)
	return r
}
