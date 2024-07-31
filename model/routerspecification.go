package model

import (
	"github.com/gofiber/fiber/v2"
)

type RouterSpecification struct {
	Api                 *fiber.App
	Logger              Logger
	ProxyRouteQR        string
	ProxyRouteStatistic string
}
