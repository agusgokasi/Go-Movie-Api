package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ok

func Ok(c *gin.Context, data interface{}) {
	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func OkWithMessage(c *gin.Context, message interface{}) {
	response := Response{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("%v", message),
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}

func NoContent(c *gin.Context) {
	response := Response{
		Status:  http.StatusNoContent,
		Message: "No Content",
		Data:    nil,
	}
	c.JSON(http.StatusNoContent, response)
}

// not ok

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	response := Response{
		Status:  http.StatusBadRequest,
		Message: message,
		Data:    nil,
	}
	if len(data) > 0 {
		response.Data = data[0]
	}
	c.JSON(http.StatusBadRequest, response)
}

func Unauthorized(c *gin.Context, message string) {
	response := Response{
		Status:  http.StatusUnauthorized,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusInternalServerError, response)
}

func NotFound(c *gin.Context, message string) {
	response := Response{
		Status:  http.StatusNotFound,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusNotFound, response)
}

func Conflict(c *gin.Context, message string) {
	response := Response{
		Status:  http.StatusConflict,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusNotFound, response)
}

func InternalServerError(c *gin.Context, message string) {
	response := Response{
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusInternalServerError, response)
}
