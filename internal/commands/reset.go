package commands

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.DataBase.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't reset the database: %w\n", err)
	}

	fmt.Println("database was reset!")

	return nil
}
