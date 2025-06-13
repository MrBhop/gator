package commands

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: %v <username>", cmd.name)
	}
	userName := cmd.args[0]

	if err := s.Config.SetUser(userName); err != nil {
		return fmt.Errorf("Couln't set user: %w", err)
	}

	fmt.Println("User name was set")
	return nil
}
