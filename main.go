package main

import (
	"fmt"
	"log"

	"github.com/MrBhop/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error  reading config: %v\n", err)
	}

	fmt.Printf("Read user config: %+v\n", cfg)

	if err := cfg.SetUser("florian"); err != nil {
		log.Fatalf("couldn't set current user: %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error  reading config: %v\n", err)
	}

	fmt.Printf("%v\n", cfg)
}
