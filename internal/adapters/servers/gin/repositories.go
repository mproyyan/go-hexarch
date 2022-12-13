package ginserver

import (
	"log"

	rmy "github.com/mproyyan/gin-rest-api/internal/application/repositories/mysql"
)

func (gs *GinServer) registerRepositories(datasource any) {
	gs.repositories["productRepository"] = rmy.NewProductRepository()
}

func (gs *GinServer) findRepository(name string) any {
	repository, found := gs.repositories[name]
	if !found {
		log.Fatalf("cannot find repository with name %s", name)
	}

	return repository
}
