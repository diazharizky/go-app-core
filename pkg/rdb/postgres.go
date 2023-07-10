package rdb

import (
	"fmt"

	"github.com/diazharizky/go-app-core/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

type postgresClient struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
}

func init() {
	config.Global.SetDefault("rdb.postgres.host", "localhost")
	config.Global.SetDefault("rdb.postgres.port", 5432)
	config.Global.SetDefault("rdb.postgres.user", "goappcore")
	config.Global.SetDefault("rdb.postgres.password", "goappcore")
	config.Global.SetDefault("rdb.postgres.dbName", "goappcore")
	config.Global.SetDefault("rdb.postgres.sslMode", "disable")
}

func NewPostgresClient() postgresClient {
	return postgresClient{
		host:     config.Global.GetString("rdb.postgres.host"),
		port:     config.Global.GetInt("rdb.postgres.port"),
		user:     config.Global.GetString("rdb.postgres.user"),
		password: config.Global.GetString("rdb.postgres.password"),
		dbName:   config.Global.GetString("rdb.postgres.dbName"),
		sslMode:  config.Global.GetString("rdb.postgres.sslMode"),
	}
}

func (client postgresClient) GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(client.dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.Use(tracing.NewPlugin()); err != nil {
		return nil, err
	}

	return db, nil
}

func (client postgresClient) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		client.host,
		client.port,
		client.user,
		client.password,
		client.dbName,
		client.sslMode,
	)
}
