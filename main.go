package main

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maliByatzes/parte-server/api"
	"github.com/maliByatzes/parte-server/config"
	db "github.com/maliByatzes/parte-server/db/sqlc"
	"github.com/maliByatzes/parte-server/worker"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	connPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.HostName,
		config.Database.Port,
		config.Database.Database,
	))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create db pool")
	}

	store := db.NewStore(connPool)

	// TODO: Run DB migration here

	redisOpt := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", config.Cache.HostName, config.Cache.Port),
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	go runTaskProcessor(redisOpt, store)
	runGinServer(config, store, taskDistributor)
}

func runGinServer(config config.Config, store db.Store, taskDIstributor worker.TaskDistributor) {
	server, err := api.NewServer(config, store, taskDIstributor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(fmt.Sprintf("%s:%d", config.HTTP.HostName, config.HTTP.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func runTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) {
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store)
	log.Info().Msg("start processor task")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
