package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maliByatzes/parte-server/api"
	db "github.com/maliByatzes/parte-server/db/sqlc"
	"github.com/maliByatzes/parte-server/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load in config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal("cannot create db pool:", err)
	}

	store := db.NewStore(connPool)

	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.HttpAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
