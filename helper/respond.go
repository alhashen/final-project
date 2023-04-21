package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var (
	ErrBadRequest      = errors.New("invalid request")
	ErrNotFound        = errors.New("record not found")
	ErrInvalidAuth     = errors.New("Invalid email or password")
	ErrForbiddenAccess = errors.New("Forbidden access")
)

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "",
		Data:    data,
	})
}

func OkWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    nil,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, Response{
		Status:  http.StatusForbidden,
		Message: "Forbidden: " + ErrForbiddenAccess.Error(),
		Data:    nil,
	})
}

func Unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized: " + ErrInvalidAuth.Error(),
		Data:    nil,
	})
}

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	/* obj := gin.H{"status": http.StatusBadRequest, "message": message}
	if len(data) > 0 {
		obj["data"] = data[0]
	} */
	c.JSON(http.StatusBadRequest, Response{
		Status:  http.StatusBadRequest,
		Message: message,
		Data:    nil,
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  http.StatusNotFound,
		Message: message,
		Data:    nil,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
}

func StatusCreated(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Response{
		Status:  http.StatusCreated,
		Message: "",
		Data:    data,
	})
}
