package bootstrap

import "os"

const _nameAppDefault = "interseguro-challenge-gateway"

func getApplicationName() string {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return _nameAppDefault
	}

	return appName
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}

func getAllowedOrigins() string {
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		return "*"
	}

	return allowedOrigins
}

func getAllowedMethods() string {
	allowedMethods := os.Getenv("ALLOWED_METHODS")
	if allowedMethods == "" {
		return "GET,POST"
	}

	return allowedMethods
}
