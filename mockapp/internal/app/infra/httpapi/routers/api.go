package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/infra/httpapi/controllers"
)

type ApiRouter struct {
	controller *controllers.ApiController
}

func (a *ApiRouter) Load(r *fiber.App) {
	r.Get("/apis", a.controller.Get)
}

func NewApiRouter(
	controller *controllers.ApiController,
) *ApiRouter {
	return &ApiRouter{
		controller: controller,
	}
}
