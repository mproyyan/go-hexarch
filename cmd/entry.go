package cmd

import (
	"github.com/mproyyan/gin-rest-api/env"
	gs "github.com/mproyyan/gin-rest-api/internal/adapters/servers/gin"
	"github.com/mproyyan/gin-rest-api/internal/ports"
)

func Boot() ports.Server {
	// load environment
	config := env.Environment{}
	env.Load(".env", &config)

	server := gs.NewGinServer(config.APP_PORT)
	return server
}
