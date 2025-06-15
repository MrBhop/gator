package commands

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %v <username>\n", cmd.name)
	}
	userName := cmd.args[0]

	return login(s, userName)
}

func login(s *state, userName string) error {
	if _, err := s.DataBase.GetUser(context.Background(), userName); err != nil {
		return fmt.Errorf("User with name '%v' does not exist\n", userName)
	}

	if err := s.Config.SetUser(userName); err != nil {
		return fmt.Errorf("Couldn't set user: %w\n", err)
	}

	fmt.Println("User name was set")
	return nil
}
