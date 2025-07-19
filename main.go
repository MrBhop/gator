package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/MrBhop/gator/internal/commands"
	"github.com/MrBhop/gator/internal/config"
	"github.com/MrBhop/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to the database: %v\n", err)
	}
	defer db.Close()

	state := commands.NewState(&cfg, database.New(db))
	commandList := commands.GetCommands()
	cmd := commands.NewCommand(os.Args[1], os.Args[2:])

	if err := commandList.Run(state, cmd); err != nil {
		log.Fatal(err)
	}
}
