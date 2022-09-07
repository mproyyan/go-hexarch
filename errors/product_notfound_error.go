package errors

import "github.com/gin-gonic/gin"

type ProductNotFoundErr struct {
	StatusCode int    `json:"status"`
	Type       string `json:"type"`
	Title      string `json:"title"`
	Detail     string `json:"detail"`
}

func NewProductNotFoundErr() *ProductNotFoundErr {
	return &ProductNotFoundErr{
		StatusCode: 404,
		Type:       "ProductNotFoundErr",
		Title:      "Product not found.",
	}
}

func ProductNotFoundErrHandler(c *gin.Context, err error) {
	pne := err.(*ProductNotFoundErr)
	c.JSON(pne.StatusCode, pne)
	c.Abort()
}

func (pne *ProductNotFoundErr) Wrap(err error) *ProductNotFoundErr {
	pne.Detail = err.Error()
	return pne
}

func (pne *ProductNotFoundErr) Error() string {
	return pne.Detail
}
