package main

import (
	"book-rest-api-showcase/app"
	"book-rest-api-showcase/config"

	_ "book-rest-api-showcase/docs"
)

// @title Book RESTFull-API
// @Version 1.0
// @description A Book RESTFull API using Gin + GORM

// @host 0.0.0.0:8081
// @BasePath /
// @securityDefinitions.basic  BasicAuth

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	err = config.InitGorm()
	if err != nil {
		panic(err)
	}

}

func main() {
	app.StartApplication()
}
