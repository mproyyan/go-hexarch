package ginserver

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mproyyan/gin-rest-api/middlewares"
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

	gs.configure()

	gs.registerRepositories(datasource)
	gs.registerServices(datasource)
	gs.registerHandlers()
	gs.registerRoutes()

	return gs
}

func (gs *GinServer) Run() {
	log.Fatal(
		gs.Engine.Run(fmt.Sprintf(":%s", gs.port)),
	)
}

func (gs *GinServer) configure() {
	// configure field name when validation error occured
	// so this config will use json tag name for field name instead of
	// using struct field name
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	// register global middleware for error handling
	gs.Engine.Use(middlewares.ErrorHandler())
}
