package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maliByatzes/parte-server/api"
	"github.com/maliByatzes/parte-server/config"
	db "github.com/maliByatzes/parte-server/db/sqlc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load in config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.HostName,
		config.Database.Port,
		config.Database.Database,
	))
	if err != nil {
		log.Fatal("cannot create db pool:", err)
	}

	store := db.NewStore(connPool)

	runGinServer(config, store)
}

func runGinServer(config config.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(fmt.Sprintf("%s:%d", config.HTTP.HostName, config.HTTP.Port))
	if err != nil {
		log.Fatal("cannot start server")
	}
}
