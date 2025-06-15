package commands

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {
	users, err := s.DataBase.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't list users: %w\n", err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf(" - %v", user.Name)
		if user.Name == s.Config.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Print("\n")
	}

	return nil
}
