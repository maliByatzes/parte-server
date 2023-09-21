# migrateschema:
# 	migrate create -ext sql -dir db/migration -seq init_<name>

# Create a docker container for postgres
postgres:
	docker run --name postgres1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Maliborh521908 -d postgres:16-alpine

# Create a database in the docker container
createdb:
	docker exec -it postgres1 createdb --username=root --owner=root partedb 

# Delete a database from the docker container
dropdb:
	docker exec -it postgres1 dropdb partedb

# Migrate up
migup:
	migrate -path db/migration -database "postgresql://root:Maliborh521908@localhost:5432/partedb?sslmode=disable" -verbose up

# Migrate down
migdown:
	migrate -path db/migration -database "postgresql://root:Maliborh521908@localhost:5432/partedb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migup migdown
