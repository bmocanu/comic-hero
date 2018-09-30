package main

import (
	"comic-hero/config"
	"fmt"
)

func main() {
	config := config.LoadConfiguration("config/comic-hero.json")
	fmt.Println(config.Server.Address)
	fmt.Println(config.Server.Port)
	fmt.Println(config.Server.ContextPath)
}
