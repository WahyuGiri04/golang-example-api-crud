package util

import (
	"example-api/model/base"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func BindJSONGeneric[T any](c *gin.Context, obj *T) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		response := base.BaseResponse{
			Status:  Failed,
			Message: "Invalid request : " + err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return false
	}
	// Validasi struct
	if err := validate.Struct(obj); err != nil {
		response := base.BaseResponse{
			Status:  Failed,
			Message: "Validation failed : " + err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return false
	}
	return true
}
