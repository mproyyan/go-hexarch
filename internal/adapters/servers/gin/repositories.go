package ginserver

import (
	"database/sql"
	"log"

	"github.com/mproyyan/gin-rest-api/helpers"
	rmy "github.com/mproyyan/gin-rest-api/internal/application/repositories/mysql"
)

func (gs *GinServer) registerRepositories(datasource any) {
	sqldb := helpers.GetDB[*sql.DB](datasource)

	gs.repositories["productRepository"] = rmy.NewProductRepository(sqldb)
}

func (gs *GinServer) findRepository(name string) any {
	repository, found := gs.repositories[name]
	if !found {
		log.Fatalf("cannot find repository with name %s", name)
	}

	return repository
}
