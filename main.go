package main

import (
	"fmt"

	"github.com/TeodorStamenov/outdoorsy/util"
)

func main() {

	config, err := util.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", config)

}
