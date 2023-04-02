package main

import (
	"book-rest-api-showcase/app"
	"book-rest-api-showcase/config"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = config.InitGorm()
	if err != nil {
		panic(err)
	}

}

func main() {
	app.StartApplication()
}
