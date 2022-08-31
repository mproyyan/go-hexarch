package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mproyyan/gin-rest-api/env"
)

type MySql struct {
	config env.Environment
}

func NewMysqlDB(config env.Environment) *MySql {
	return &MySql{
		config: config,
	}
}

func (m *MySql) ConnectDB() *sql.DB {
	db, err := sql.Open(
		m.config.DB_CONN,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			m.config.DB_USERNAME,
			m.config.DB_PASSWORD,
			m.config.DB_HOST,
			m.config.DB_PORT,
			m.config.DB_NAME,
		))

	if err != nil {
		log.Fatal(err)
	}

	return db
}
