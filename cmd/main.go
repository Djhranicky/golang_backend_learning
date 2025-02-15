// Following along with https://youtu.be/7VLmLOiQ3ck?si=pTFmzFAwYPZJb-JE&t=1126
// Could do this one as well https://www.youtube.com/watch?v=h3fqD6IprIA

package main

import (
	"database/sql"
	"log"

	"github.com/djhranicky/golang_backend_learning/cmd/api"
	"github.com/djhranicky/golang_backend_learning/config"
	"github.com/djhranicky/golang_backend_learning/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
