package app

import (
	"book-rest-api-showcase/config"
	"book-rest-api-showcase/handler"
	"book-rest-api-showcase/repository"
	"book-rest-api-showcase/route"
	"book-rest-api-showcase/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))

}
