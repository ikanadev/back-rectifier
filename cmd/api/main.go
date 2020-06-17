package main

import (
	"fmt"
	"os"

	"github.com/vkevv/back-rectifier/pkg/api"
	"github.com/vkevv/back-rectifier/pkg/config"
)

func main() {
	projectDir, err := os.Getwd()
	checkError(err)
	conf, err := config.GetConfig(projectDir + "/cmd/api/config.yaml")
	checkError(err)
	checkError(api.StartAPI(conf))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR", err.Error())
		panic(err.Error())
	}
}
