package commands

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.DataBase.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("the following error occured, trying to rest the database: %w", err)
	}

	fmt.Print("database was reset")

	return nil
}
