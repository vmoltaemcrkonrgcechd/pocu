package main

import (
	"github.com/vmoltaemcrkonrgcechd/pocu/config"
	"github.com/vmoltaemcrkonrgcechd/pocu/internal/app"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
