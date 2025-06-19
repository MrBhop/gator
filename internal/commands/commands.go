package commands

import (
	"context"
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

type handlerFunction = func(*state, command) error

type commands struct {
	registeredCommands map[string] handlerFunction
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
	newCommandList.register("users", handlerUsers)
	newCommandList.register("agg", handlerAgg)
	newCommandList.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	newCommandList.register("feeds", handlerFeeds)
	newCommandList.register("follow", middlewareLoggedIn(handlerFollow))
	newCommandList.register("following", middlewareLoggedIn(handlerFollowing))
	newCommandList.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	commandList = &newCommandList
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) Run(s *state, cmd command) error {
	callback, exists := c.registeredCommands[cmd.name]
	if !exists {
		return fmt.Errorf("command '%v' does not exist\n", cmd.name)
	}

	return callback(s, cmd)
}

type middlewareHandlerFunction = func(s *state, cmd command, user database.User) error

func middlewareLoggedIn(handler middlewareHandlerFunction) handlerFunction {
	return func(s *state, cmd command) error {
		user, err := s.DataBase.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't fetch current user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
