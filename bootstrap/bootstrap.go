package bootstrap

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/MikelSot/interseguro-challenge-gateway/infrastructure/handler"
	"github.com/MikelSot/interseguro-challenge-gateway/infrastructure/handler/response"
	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

func Run() {
	_ = godotenv.Load()

	app := newFiber(response.ErrorHandler)
	logger := newLogger(false)

	handler.InitRoutes(model.RouterSpecification{
		Api:                 app,
		Logger:              logger,
		ProxyRouteQR:        getProxyRouteQR(),
		ProxyRouteStatistic: getProxyRouteStatistic(),
	})

	log.Fatal(app.Listen(getPort()))
}
