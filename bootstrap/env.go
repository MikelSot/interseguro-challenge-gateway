package bootstrap

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const _nameAppDefault = "interseguro-challenge-gateway"

const _portDefault = ":8080"

const (
	_allowOriginsDefault = "*"
	_allowMethodsDefault = "GET,POST"
)

const (
	_proxyRouteQRDefault        = "/qr"
	_proxyRouteStatisticDefault = "/statistic"
	_proxyRouteAuthDefault      = "/auth"
)

func getApplicationName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}

func getPort() string {
	port := os.Getenv("FIBER_PORT")
	if port == "" {
		return _portDefault
	}

	return port
}

func getAllowOrigins() string {
	allowedOrigins := os.Getenv("ALLOW_ORIGINS")
	if allowedOrigins == "" {
		return _allowOriginsDefault
	}

	return allowedOrigins
}

func getAllowMethods() string {
	allowedMethods := os.Getenv("ALLOW_METHODS")
	if allowedMethods == "" {
		return _allowMethodsDefault
	}

	return allowedMethods
}

func getProxyRouteQR() string {
	proxyRouteQR := os.Getenv("PROXY_ROUTE_QR")
	if proxyRouteQR == "" {
		log.Warn("proxy route qr not found")

		return _proxyRouteQRDefault
	}

	return proxyRouteQR
}

func getProxyRouteStatistic() string {
	proxyRouteStatistic := os.Getenv("PROXY_ROUTE_STATISTIC")
	if proxyRouteStatistic == "" {
		log.Warn("proxy route statistic not found")

		return _proxyRouteStatisticDefault
	}

	return proxyRouteStatistic
}

func getProxyRouteAuth() string {
	proxyRouteAuth := os.Getenv("PROXY_ROUTE_AUTH")
	if proxyRouteAuth == "" {
		log.Warn("proxy route auth not found")

		return _proxyRouteAuthDefault
	}

	return proxyRouteAuth
}

func getSignKey() string {
	signKey := os.Getenv("SIGN_KEY")
	if signKey == "" {
		log.Warn("sign key not found")

		return ""
	}

	return signKey
}
