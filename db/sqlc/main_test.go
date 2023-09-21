package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), "postgresql://root:Maliborh521908@localhost:5432/partedb?sslmode=disable")
	if err != nil {
		log.Fatal("cannot to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
