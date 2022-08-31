package ginserver

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	*gin.Engine
	port         string
	repositories map[string]interface{}
	services     map[string]interface{}
	handlers     map[string]interface{}
}

func NewGinServer(datasource interface{}, port string) *GinServer {
	gs := &GinServer{
		Engine:       gin.Default(),
		port:         port,
		repositories: make(map[string]interface{}),
		services:     make(map[string]interface{}),
		handlers:     make(map[string]interface{}),
	}

	return gs
}

func (gs *GinServer) Run() {
	log.Fatal(
		gs.Engine.Run(fmt.Sprintf(":%s", gs.port)),
	)
}
