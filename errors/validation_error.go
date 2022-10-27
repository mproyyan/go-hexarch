package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	tl "github.com/mproyyan/validation-error-translator"
)

type ErrorField struct {
	FieldName string `json:"field"`
	Reason    string `json:"reason"`
}

type ValidationErr struct {
	StatusCode int          `json:"status"`
	Type       string       `json:"type"`
	Title      string       `json:"title"`
	Detail     string       `json:"detail"`
	Errors     []ErrorField `json:"errors"`
}

func ValidationErrHandler(c *gin.Context, err error) {
	verr := err.(validator.ValidationErrors)
	errorFields, totalErrorFields := extractErrorField(verr)
	var detail string
	if totalErrorFields > 1 {
		detail = fmt.Sprintf("there are %d fields where an error occurs", totalErrorFields)
	} else {
		detail = fmt.Sprintf("there are %d field where an error occurs", totalErrorFields)
	}

	valErr := &ValidationErr{
		StatusCode: 400,
		Type:       "ValidationErr",
		Title:      "Validation error.",
		Detail:     detail,
		Errors:     errorFields,
	}

	c.JSON(valErr.StatusCode, valErr)
	c.Abort()
}

func extractErrorField(verr validator.ValidationErrors) ([]ErrorField, int) {
	var errorFields []ErrorField
	for _, field := range verr {
		errField := ErrorField{
			FieldName: field.Field(),
			Reason:    tl.Translate(field),
		}
		errorFields = append(errorFields, errField)
	}

	return errorFields, len(verr)
}
