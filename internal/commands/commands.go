package commands

import (
	"fmt"

	"github.com/MrBhop/BlogAggregator/internal/config"
	"github.com/MrBhop/BlogAggregator/internal/database"
)

type state struct {
	Config *config.Config
	DataBase *database.Queries
}

func NewState(config *config.Config, database *database.Queries) *state {
	return &state{
		Config: config,
		DataBase: database,
	}
}

type command struct {
	name string
	args []string
}

func NewCommand(name string, args []string) command {
	return command{
		name: name,
		args: args,
	}
}

type commands struct {
	registeredCommands map[string] func(*state, command) error
}

var commandList *commands

func GetCommands() *commands {
	if commandList == nil {
		initializeCommandList()
	}

	return commandList
}

func initializeCommandList() {
	newCommandList := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	newCommandList.register("login", handlerLogin)
	newCommandList.register("register", handlerRegister)
	newCommandList.register("reset", handlerReset)

	commandList = &newCommandList
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) Run(s *state, cmd command) error {
	callback, exists := c.registeredCommands[cmd.name]
	if !exists {
		return fmt.Errorf("command '%v' does not exist", cmd.name)
	}

	return callback(s, cmd)
}
