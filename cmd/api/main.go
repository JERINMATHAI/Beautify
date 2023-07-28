package main

import (
	"beautify/cmd/api/docs"
	"beautify/pkg/config"
	"beautify/pkg/di"
	"log"
)

//	@SecurityDefinition	BearerAuth
//	@TokenUrl			/auth/token

//	@securityDefinitions.Bearer		type apiKey
//	@securityDefinitions.Bearer		name Authorization
//	@securityDefinitions.Bearer		in header
//	@securityDefinitions.BasicAuth	type basic

func main() {

	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "BEAUTIFY"
	docs.SwaggerInfo.Description = "This is an online store for purchasing high quality beauty products"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Cannot load config: ", configErr)
	}
	//Initialize API server
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("Cannot start server: ", diErr)
	} else {
		server.Start()
	}

}
