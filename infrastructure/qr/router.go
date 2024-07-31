package qr

import (
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

const (
	_privateRoutePrefix = "/api/v1/qr"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	privateRoutes(spec.Api, handler)
}

func buildHandler(spec model.RouterSpecification) handler {
	return NewHandler(spec.ProxyRouteQR)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.All("", handler.Proxy)
}
