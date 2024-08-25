package db

import (
    "context"
    "log"
    "os"
    "testing"

    "github.com/jackc/pgx/v5/pgxpool"
)

const (
    dbSource = "postgresql://root:root@localhost:5432/bank_project?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
    conn, err := pgxpool.New(context.Background(), dbSource)
    if err != nil {
        log.Fatal("Cannot connect to database:", err)
    }
    defer conn.Close()

    testQueries = New(conn)
    os.Exit(m.Run())
}