.PHONY: up
up:
	docker compose up -d
	sleep 3

.PHONY: reset
reset:
	docker compose down
	make up

.PHONY: db
db:
	docker exec -it parte_db psql postgresql://admin:admin@localhost:5432/app

.PHONY: migup
migup:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/app?sslmode=disable" -verbose up

.PHONY: migdown
migdown:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/app?sslmode=disable" -verbose down

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test
test:
	make migdown
	make migup
	go test -count=1 -p 1 ./... -v

.PHONY: server
server:
	go run main.go

