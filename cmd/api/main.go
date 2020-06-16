package main

import (
	"fmt"

	"github.com/vkevv/back-rectifier/pkg/api"
	"github.com/vkevv/back-rectifier/pkg/config"
)

func main() {
	conf, err := config.GetConfig("./cmd/api/config.yaml")
	checkError(err)
	checkError(api.StartAPI(conf))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR", err.Error())
		panic(err.Error())
	}
}
