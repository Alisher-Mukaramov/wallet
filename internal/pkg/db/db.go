package db

import (
	"fmt"
	cfg "github.com/Alisher-Mukaramov/wallet/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"log"
)

var (
	Module           = fx.Provide(newPostgres)
	config           = cfg.GetConfig()
	dbHost           = config.Database.Host
	dbPort           = config.Database.Port
	dbName           = config.Database.Name
	dbUser           = config.Database.Username
	dbPassword       = config.Database.Password
	connectionString = fmt.Sprintf("host=%s port=%s dbname='%s' user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)
)

type Idb interface {
	DBInstance() *sqlx.DB
}

type db struct {
	postgres *sqlx.DB
}

func newPostgres() Idb {
	postgresDB, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	return &db{postgres: postgresDB}
}

func (d *db) DBInstance() *sqlx.DB {
	return d.postgres
}
