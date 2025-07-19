package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/MrBhop/gator/internal/database"
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

	fmt.Printf("followed feed: \n%v\n", CreateFeedFollowRowToString(followEntry))
	return nil
}


func CreateFeedFollowRowToString(item database.CreateFeedFollowRow) string {
	output := fmt.Sprintln("CreateFeedFollowRow{")
	output += fmt.Sprintf("ID: %v\n", item.ID)
	output += fmt.Sprintf("CreatedAt: %v\n", item.CreatedAt)
	output += fmt.Sprintf("UpdatedAt: %v\n", item.UpdatedAt)
	output += fmt.Sprintf("UserID: %v\n", item.UserID)
	output += fmt.Sprintf("FeedID: %v\n", item.FeedID)
	output += fmt.Sprintf("UserName: %v\n", item.UserName)
	output += fmt.Sprintf("FeedName: %v\n", item.FeedName)
	output += fmt.Sprintln("}")
	return output
}
