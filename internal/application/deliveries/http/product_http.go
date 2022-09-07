package httpdelivery

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	cuserr "github.com/mproyyan/gin-rest-api/errors"
	"github.com/mproyyan/gin-rest-api/internal/application/domain"
)

type ProductHttp struct {
	ProductService domain.ProductService
}

func NewProductHttp(productService domain.ProductService) *ProductHttp {
	return &ProductHttp{
		ProductService: productService,
	}
}

func (ph *ProductHttp) FindAll(c *gin.Context) {
	products, err := ph.ProductService.FindAll(c.Request.Context())
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(200, products)
}

func (ph *ProductHttp) Create(c *gin.Context) {
	var request domain.ProductCreateRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	createdProduct, _ := ph.ProductService.Create(c.Request.Context(), request)
	c.JSON(201, createdProduct)
}

func (ph *ProductHttp) Find(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		notFound := fmt.Errorf("you tried to search for a product with id %s, and no results were found", productId)
		c.Error(cuserr.NewProductNotFoundErr().Wrap(notFound))
		c.Abort()
		return
	}

	product, err := ph.ProductService.Find(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(200, product)
}

func (ph *ProductHttp) Update(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		notFound := fmt.Errorf("you tried to search for a product with id %s, and no results were found", productId)
		c.Error(cuserr.NewProductNotFoundErr().Wrap(notFound))
		c.Abort()
		return
	}

	var request domain.ProductUpdateRequest
	if err = c.ShouldBind(&request); err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	request.ID = id

	product, err := ph.ProductService.Update(c.Request.Context(), request)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(200, product)
}

func (ph *ProductHttp) Delete(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		notFound := fmt.Errorf("you tried to search for a product with id %s, and no results were found", productId)
		c.Error(cuserr.NewProductNotFoundErr().Wrap(notFound))
		c.Abort()
		return
	}

	err = ph.ProductService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "Product deleted successfully",
	})
}
