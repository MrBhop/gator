package commands

import (
	"fmt"

	"github.com/MrBhop/BlogAggregator/internal/config"
)

type state struct {
	Config *config.Config
}

func NewState(config *config.Config) *state {
	return &state{
		Config: config,
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

func GetCommands() *commands {
	newCommands := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	newCommands.register("login", handlerLogin)

	return &newCommands
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
