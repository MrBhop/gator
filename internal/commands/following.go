package commands

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, _ command) error {
	user, err := s.DataBase.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't fetch user: %w", err)
	}

	followedFeeds, err := s.DataBase.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch followed feeds: %w", err)
	}

	fmt.Println("Followed Feeds:")
	for _, item := range followedFeeds {
		fmt.Printf("  - %v\n", item)
	}

	return nil
}
