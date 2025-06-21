package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/MrBhop/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: %v <url>\n", cmd.name)
	}
	url := cmd.args[0]

	feed, err := s.DataBase.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("the feed you are trying to follow does not exist yet: %w\n", err)
	}

	return follow(s, user.ID, feed.ID)
}


func follow(s *state, userId uuid.UUID, feedId uuid.UUID) error {
	followEntry, err := s.DataBase.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: userId,
		FeedID: feedId,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("followed feed: \n%+v\n", followEntry)
	return nil
}
