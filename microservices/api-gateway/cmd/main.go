package main

import (
	"fmt"
	"log"

	"github.com/hexhoc/api-gateway/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	fmt.Println("Starting api gateways")
	fmt.Println(config)
}
