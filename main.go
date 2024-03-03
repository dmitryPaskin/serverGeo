package main

import (
	"GeoAPI/router"
)

// @title My API
// @version 1.1
// @description This is a sample API for address searching and geocoding using Dadata API.
// @host localhost:8080
// @termsOfService http://localhost:8080/swagger/index.html
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := router.NewRouter("/api/address/geocode", "/api/address/search")
	r.StartRouter()
}
