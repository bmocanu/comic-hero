package main

import (
	"comic-hero/config"
	"fmt"
	"github.com/mileusna/crontab"
	"log"
	_ "github.com/mileusna/crontab"
	"time"
)

func main() {
	cfg, err := config.LoadConfiguration("config/comic-hero.json")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg.Server.Address)
	fmt.Println(cfg.Server.Port)
	fmt.Println(cfg.Server.ContextPath)
	fmt.Println(cfg.Comics[0].Name)

	ctab := crontab.New()
	ctab.MustAddJob("* * * * *", doTheTask)
	time.Sleep(10 * time.Minute)
}

func doTheTask() {
	log.Println("Inside the task!")
}
