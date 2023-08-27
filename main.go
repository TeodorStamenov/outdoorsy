package main

import (
	"fmt"

	"github.com/TeodorStamenov/outdoorsy/db"
	"github.com/TeodorStamenov/outdoorsy/util"
)

func main() {

	config, err := util.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", config)

	conn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Db.Addr, config.Db.Port, config.Db.User, config.Db.Pass, config.Db.Name)

	fmt.Printf("Connection string: %s\n", conn)
	db.Connect(conn)
}
