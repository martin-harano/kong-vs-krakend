package main

import (
	"github.com/google/uuid"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/domain/faker"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/infra/httpapi"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/infra/httpapi/controllers"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/infra/httpapi/routers"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func main() {
	uuid.EnableRandPool()

	app := fx.New(
		controllers.Module,
		routers.Module,
		httpapi.Module,
		faker.Module,
		fx.Invoke(func(*fasthttp.Server) {}),
	)

	app.Run()
}
