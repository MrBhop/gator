package commands

import (
	"context"
	"fmt"

	"github.com/MrBhop/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: %v <feed_url>\n", cmd.name)
	}
	url := cmd.args[0]

	feed, err := s.DataBase.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w\n", err)
	}

	if err := s.DataBase.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("couldn't delete user: %w\n", err)
	}

	fmt.Printf("user '%v' unfollowed feed '%v'\n", user.Name, feed.Name)
	return nil
}
