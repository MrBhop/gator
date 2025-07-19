package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MrBhop/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) >= 1 {
		if val, err := strconv.Atoi(cmd.args[0]); err == nil {
			limit = val
		}
	}

	posts, err := s.DataBase.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't fetch posts for user: %w", err)
	}

	fmt.Println("Listing posts:")
	for i, item := range posts {
		fmt.Printf("Post %v:\n", i + 1)
		fmt.Printf("%+v\n", item)
		fmt.Println()
	}

	return nil
}
