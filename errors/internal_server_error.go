package errors

import "github.com/gin-gonic/gin"

type InternalServerErr struct {
	StatusCode int    `json:"status"`
	Type       string `json:"type"`
	Title      string `json:"title"`
	Detail     string `json:"detail"`
}

func NewInternalServerErr() *InternalServerErr {
	return &InternalServerErr{
		StatusCode: 500,
		Type:       "InternalServerErr",
		Title:      "Internal server error.",
	}
}

func InternalServerErrHandler(c *gin.Context, err error) {
	ise := err.(*InternalServerErr)
	c.JSON(ise.StatusCode, ise)
	c.Abort()
}

func (ise *InternalServerErr) Wrap(err error) *InternalServerErr {
	ise.Detail = err.Error()
	return ise
}

func (ise *InternalServerErr) Error() string {
	return ise.Detail
}
