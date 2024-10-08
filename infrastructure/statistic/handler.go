package statistic

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

type handler struct {
	proxyRouterStatistic string
}

func newHandler(proxyRouterStatistic string) handler {
	return handler{proxyRouterStatistic}
}

func (h handler) Proxy(c *fiber.Ctx) error {
	url := fmt.Sprintf("%s%s", h.proxyRouterStatistic, c.OriginalURL())
	if err := proxy.Do(c, url); err != nil {
		log.Warn("¡Uy! Fallo al enviar la solicitud a la API Statistic", err.Error())

		return c.Status(fiber.StatusBadGateway).JSON(model.MessageResponse{
			Errors: model.Responses{
				{Code: model.UnexpectedError, Message: "¡Uy! Fallo al enviar la solicitud a la API Statistic"},
			},
		})
	}
	c.Response().Header.Del(fiber.HeaderServer)

	return nil
}
