package main

import (
	"fmt"

	"github.com/TeodorStamenov/outdoorsy/db"
	"github.com/TeodorStamenov/outdoorsy/service"
	"github.com/TeodorStamenov/outdoorsy/util"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := db.NewDb(config)
	if err != nil {
		panic(err)
	}

	app, err := service.NewService(config, db)
	if err != nil {
		panic(err)
	}

	app.Run()
}
