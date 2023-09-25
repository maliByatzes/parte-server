package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/maliByatzes/parte-server/db/sqlc"
)

type TaskProcessor interface {
	Start() error
	ProcessorTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessorTaskSendVerifyEmail)

	return processor.server.Start(mux)
}
