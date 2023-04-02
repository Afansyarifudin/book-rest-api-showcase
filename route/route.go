package route

import (
	"book-rest-api-showcase/docs"
	"book-rest-api-showcase/handler"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

//	@title			Book RESTFul API
//	@version		1.0
//	@description	This is a Book RestFul API Using Golang.
//	@host			localhost:8081
//	@BasePath		/

func RegisterApi(r *gin.Engine, server handler.HttpServer) {

	docs.SwaggerInfo.Title = "Book RESTFul API"
	docs.SwaggerInfo.Description = "This is a Book RestFul API Using Golang."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	api := r.Group("/books")
	{
		// Create
		api.POST("", server.CreateBook)
		// Read All
		api.GET("", server.GetAllBooks)
		// Read
		api.GET("/:id", server.GetBookByID)
		// Update
		api.PUT("/:id", server.UpdateBook)
		// DELETE
		api.DELETE("/:id", server.DeleteBook)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
