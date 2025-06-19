package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/MrBhop/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Usage: %v <name> <url>", cmd.name)
	}

	name := cmd.args[0]
	url := cmd.args[1]

	user, err := s.DataBase.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't fetch current user: %w", err)
	}

	feed, err := s.DataBase.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	return follow(s, feed.UserID, feed.ID)
}
