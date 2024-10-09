package main

import "backend/internal/app"

// @title						GulTix API
// @version					2.0
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used
func main() {
	server := app.InitHttp()

	server.InitRouteAndServe()
}
