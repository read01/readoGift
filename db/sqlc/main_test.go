package db

import (
	"database/sql"
	"log"
	"os"
	"readoGift/util"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQueries *Queries
var testDB *sql.DB


func TestMain(m *testing.M) {
	config,err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()
	testQueries = New(testDB)
	os.Exit(m.Run())
}
