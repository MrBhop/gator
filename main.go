package main

import (
	"log"
	"os"

	"github.com/MrBhop/BlogAggregator/internal/config"
	"github.com/MrBhop/BlogAggregator/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	state := (commands.NewState(&cfg))
	commandList := commands.GetCommands()
	cmd := commands.NewCommand(os.Args[1], os.Args[2:])

	if err := commandList.Run(state, cmd); err != nil {
		log.Fatal(err)
	}
}
