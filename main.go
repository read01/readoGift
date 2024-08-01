package main

import (
	"database/sql"
	"log"
	"readoGift/api"
	db "readoGift/db/sqlc"
	"readoGift/util"

	_ "github.com/jackc/pgx/v5/stdlib"
)


func main() {
	config,err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:",err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ",err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server : ",err)
	}

}