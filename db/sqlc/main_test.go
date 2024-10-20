package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbDriver    string
	dbSource    string
	testQueries *Queries
)

func TestMain(m *testing.M) {

	dbDriver = os.Getenv("DB_DRIVER")
	dbSource = os.Getenv("DB_SOURCE")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
