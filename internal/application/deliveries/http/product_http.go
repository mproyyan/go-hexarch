package httpdelivery

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
	products := ph.ProductService.FindAll(c.Request.Context())
	c.JSON(200, products)
}

func (ph *ProductHttp) Create(c *gin.Context) {
	var request domain.ProductCreateRequest
	if err := c.ShouldBind(&request); err != nil {
		log.Fatal(err)
	}

	createdProduct := ph.ProductService.Create(c.Request.Context(), request)
	c.JSON(201, createdProduct)
}

func (ph *ProductHttp) Find(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(404)
	}

	product := ph.ProductService.Find(c.Request.Context(), productId)
	c.JSON(200, product)
}

func (ph *ProductHttp) Update(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(404)
	}

	var request domain.ProductUpdateRequest
	c.ShouldBind(&request)
	request.ID = productId

	product := ph.ProductService.Update(c.Request.Context(), request)
	c.JSON(200, product)
}

func (ph *ProductHttp) Delete(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(404)
	}

	ph.ProductService.Delete(c.Request.Context(), productId)
	c.JSON(200, gin.H{
		"message": "Product deleted successfully",
	})
}
