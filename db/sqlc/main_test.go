package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/maliByatzes/parte-server/config"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.HostName,
		config.Database.Port,
		config.Database.Database,
	))
	if err != nil {
		log.Fatal("cannot to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
