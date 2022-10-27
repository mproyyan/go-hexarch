package cmd

import (
	"fmt"
	"log"

	"github.com/mproyyan/gin-rest-api/env"
	"github.com/mproyyan/gin-rest-api/internal/adapters/databases"
	gs "github.com/mproyyan/gin-rest-api/internal/adapters/servers/gin"
	"github.com/mproyyan/gin-rest-api/internal/ports"
)

func Boot() ports.Server {
	// load environment
	config := env.Environment{}
	env.Load(".env", &config)

	// get database connection based on env file configuration
	datasource, err := getDB(config)
	if err != nil {
		log.Fatal(err)
	}

	server := gs.NewGinServer(datasource, config)
	return server
}

func getDB(config env.Environment) (interface{}, error) {
	switch config.DB_CONN {
	case "mysql":
		return databases.NewMysqlDB(config).ConnectDB(), nil
	default:
		return nil, fmt.Errorf("database connection for %s is not available", config.DB_CONN)
	}
}
