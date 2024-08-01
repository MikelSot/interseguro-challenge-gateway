package qr

import (
	interseguroAuth "github.com/MikelSot/interseguro-challenge-auth/bootstrap"
	"github.com/gofiber/fiber/v2"

	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

const (
	_privateRoutePrefix = "qr"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	privateRoutes(spec.Api, handler, interseguroAuth.ValidateJWT)
}

func buildHandler(spec model.RouterSpecification) handler {
	return newHandler(spec.ProxyRouteQR)
}

func privateRoutes(app *fiber.App, handler handler, middlewares ...fiber.Handler) {
	api := app.Group(_privateRoutePrefix, middlewares...)

	api.All("/api/v1/factorize", handler.Proxy)
}
