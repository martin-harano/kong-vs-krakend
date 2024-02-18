package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/infra/config"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/infra/httpapi/controllers"
)

type ApiRouter struct {
	controller *controllers.ApiController
	path       string
}

func (a *ApiRouter) Load(r *fiber.App) {
	r.Post(a.path, a.controller.Get)
}

func NewApiRouter(
	config *config.Config,
	controller *controllers.ApiController,
) *ApiRouter {
	return &ApiRouter{
		controller: controller,
		path:       config.Routers.Path,
	}
}
