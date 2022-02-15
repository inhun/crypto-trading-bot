package main

import (
	"fmt"

	"github.com/inhun/crypto-trading-bot/config"
)

func main() {
	cfg, _ := config.LoadConfig("config.json")
	fmt.Println(cfg)

	fmt.Println("hello")
}
