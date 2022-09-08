package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	cuserr "github.com/mproyyan/gin-rest-api/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch err.Err.(type) {
			case *cuserr.InternalServerErr:
				cuserr.InternalServerErrHandler(c, err.Err)
				return
			case *cuserr.ProductNotFoundErr:
				cuserr.ProductNotFoundErrHandler(c, err.Err)
				return
			case validator.ValidationErrors:
				cuserr.ValidationErrHandler(c, err.Err)
				return
			default:
				c.JSON(500, gin.H{"error": err.Err.Error()})
				c.Abort()
				return
			}
		}
	}
}
