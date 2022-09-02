package ginserver

import (
	"database/sql"
	"log"

	"github.com/mproyyan/gin-rest-api/helpers"
	rmy "github.com/mproyyan/gin-rest-api/internal/application/repositories/mysql"
	smy "github.com/mproyyan/gin-rest-api/internal/application/services/mysql"
)

func (gs *GinServer) registerServices(datasource any) {
	sqldb := helpers.GetDB[*sql.DB](datasource)
	productRepo := gs.findRepository("productRepository").(*rmy.ProductRepository)

	gs.services["productService"] = smy.NewProductService(sqldb, productRepo)
}

func (gs *GinServer) findService(name string) any {
	service, found := gs.services[name]
	if !found {
		log.Fatalf("cannot find service with name %s", name)
	}

	return service
}
