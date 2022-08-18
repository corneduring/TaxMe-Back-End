package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = ""
	dbname   = "postgres"
)

func ConnectDatabase() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password='%s' dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected!")

	return db
}

func RunScript(sql string, db *sql.DB) {
	script, err := ioutil.ReadFile(sql)
	if err != nil {
		log.Panic(err)
	}

	stringScript := string(script)
	queries := strings.Split(stringScript, ";")
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		log.Panic(err)
	}

	for i, _ := range queries {
		if _, err = tx.ExecContext(ctx, queries[i]); err != nil {
			tx.Rollback()
			fmt.Println("\n", err, "\n ....Transaction rollback!\n")
			return
		}
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("....Transaction committed\n")
	}
}
