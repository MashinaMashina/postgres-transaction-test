package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"testing"
	"time"
)

var connect *pgx.Conn
var IsoLevels = []pgx.TxIsoLevel{pgx.ReadCommitted, pgx.ReadUncommitted, pgx.Serializable, pgx.RepeatableRead}

func TestMain(m *testing.M) {
	var dsn string
	flag.StringVar(&dsn, "dsn", "", "database source name. Example: postgres://username:password@localhost:5432/database_name")

	flag.Parse()

	if dsn == "" {
		panic("not set dsn argument. Use 'go test -v -dsn dsn_value'. Example dsn_value: postgres://username:password@localhost:5432/database_name")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic(fmt.Errorf("connecting to database: %w", err))
	}

	connect = conn

	os.Exit(m.Run())
}

func Connect() *pgx.Conn {
	return connect
}
