package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrorNotFound = "record Not Found"
)

// status Ok

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func OkWithMessage(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		// "status":  http.StatusOK,
		"message": message,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

// Status Bad

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	obj := gin.H{"status": http.StatusBadRequest, "message": message}
	if len(data) > 0 {
		obj["data"] = data[0]
	}
	c.JSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": message})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": message})
}
