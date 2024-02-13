package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martin-harano/kong-vs-krakend/mockapp/internal/app/domain/benchmark"
)

type CreateApiRequest struct {
	Name    string `json:"name" validate:"required,max=32"`
	Latency int64  `json:"latency"`
}

type ApiResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ApiController struct {
	createApi *benchmark.CreateApi
}

func mapApiResponse(api *benchmark.Api) ApiResponse {
	return ApiResponse{
		ID:   api.ID,
		Name: api.Name,
	}
}

func (a *ApiController) Get(c *fiber.Ctx) error {
	var dto CreateApiRequest
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "bad request"})
	}
	api, err := a.createApi.Execute(
		dto.Name,
		dto.Latency,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.Status(fiber.StatusOK).JSON(mapApiResponse(api))
}

func NewApiController(
	createApi *benchmark.CreateApi,
) *ApiController {
	return &ApiController{
		createApi: createApi,
	}
}
