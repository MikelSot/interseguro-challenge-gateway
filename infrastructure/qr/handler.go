package qr

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"

	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

type handler struct {
	proxyRouterQR string
}

func NewHandler(proxyRouterQR string) handler {
	return handler{proxyRouterQR}
}

func (h handler) Proxy(c *fiber.Ctx) error {
	url := fmt.Sprintf("%s%s", h.proxyRouterQR, c.OriginalURL())
	if err := proxy.Do(c, url); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(model.MessageResponse{
			Errors: model.Responses{
				{Code: model.UnexpectedError, Message: "¡Uy! Fallo al enviar la solicitud a la API QR"},
			},
		})
	}
	c.Response().Header.Del(fiber.HeaderServer)

	return nil
}
