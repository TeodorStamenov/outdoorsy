package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(conn string) {
	repo, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	err = repo.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
