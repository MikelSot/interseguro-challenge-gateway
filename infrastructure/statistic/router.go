package statistic

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

const (
	_privateRoutePrefix = "/api/v1/statistic"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	privateRoutes(spec.Api, handler)
}

func buildHandler(spec model.RouterSpecification) handler {
	return NewHandler(spec.ProxyRouteStatistic)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.All("", handler.Proxy)
}
