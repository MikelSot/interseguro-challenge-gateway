package statistic

import (
	interseguroAuth "github.com/MikelSot/interseguro-challenge-auth/bootstrap"
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

const (
	_privateRoutePrefix = "statistic"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	privateRoutes(spec.Api, handler, interseguroAuth.ValidateJWT)
}

func buildHandler(spec model.RouterSpecification) handler {
	return newHandler(spec.ProxyRouteStatistic)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.All("", handler.Proxy)
}
