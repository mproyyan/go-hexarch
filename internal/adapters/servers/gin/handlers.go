package ginserver

import (
	"log"

	httpd "github.com/mproyyan/gin-rest-api/internal/application/deliveries/http"
	smy "github.com/mproyyan/gin-rest-api/internal/application/services/mysql"
)

func (gs *GinServer) registerHandlers() {
	productService := gs.findService("productService").(*smy.ProductService)

	gs.handlers["productHttp"] = httpd.NewProductHttp(productService)
}

func (gs *GinServer) findHandler(name string) any {
	handler, found := gs.handlers[name]
	if !found {
		log.Fatalf("cannot find handler with name %s", name)
	}

	return handler
}
