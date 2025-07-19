package commands

import (
	"context"
	"fmt"

	"github.com/MrBhop/gator/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {
	followedFeeds, err := s.DataBase.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch followed feeds: %w", err)
	}

	fmt.Println("Followed Feeds:")
	for _, item := range followedFeeds {
		fmt.Printf("  - %v\n", item.FeedName)
	}

	return nil
}
