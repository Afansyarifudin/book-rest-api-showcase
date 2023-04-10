package route

import (
	"book-rest-api-showcase/handler"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {

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

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
