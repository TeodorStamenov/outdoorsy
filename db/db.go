package db

import (
	"database/sql"
	"fmt"

	"github.com/TeodorStamenov/outdoorsy/util"
	_ "github.com/lib/pq"
)

type Db interface {
}

type Postgres struct {
	db *sql.DB
}

func NewDb(c util.Config) (Db, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Db.Addr, c.Db.Port, c.Db.User, c.Db.Pass, c.Db.Name)

	fmt.Printf("Database trying to connect with: %s\n", conn)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return Postgres{
		db: db,
	}, err
}
