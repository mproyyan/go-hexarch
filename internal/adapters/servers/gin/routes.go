package ginserver

import (
	"github.com/gin-gonic/gin"
	httpd "github.com/mproyyan/gin-rest-api/internal/application/deliveries/http"
)

func (gs *GinServer) productRoutes(r *gin.Engine) {
	productHttp := gs.findHandler("productHttp").(*httpd.ProductHttp)

	pr := r.Group("/api/products")
	pr.GET("/", productHttp.FindAll)
	pr.POST("/", productHttp.Create)
	pr.GET("/:id", productHttp.Find)
	pr.PUT("/:id", productHttp.Update)
	pr.DELETE("/:id", productHttp.Delete)
}

func (gs *GinServer) registerRoutes() {
	gs.productRoutes(gs.Engine)
}
