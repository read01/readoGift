package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("pgx", dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()
	testQueries = New(testDB)
	os.Exit(m.Run())
}
