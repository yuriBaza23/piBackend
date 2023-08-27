package main

import (
	"pi/cmd/server"
	"pi/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server.HttpInit(config.GetAPIConfig())
}
