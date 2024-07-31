package handler

import (
	"github.com/MikelSot/interseguro-challenge-gateway/infrastructure/qr"
	"github.com/MikelSot/interseguro-challenge-gateway/infrastructure/statistic"
	"github.com/MikelSot/interseguro-challenge-gateway/model"
)

func InitRoutes(spec model.RouterSpecification) {
	// Q
	qr.NewRouter(spec)

	// S
	statistic.NewRouter(spec)
}
