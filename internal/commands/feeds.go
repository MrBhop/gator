package commands

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.DataBase.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println("{")
		fmt.Printf("  name: %v\n", feed.Name)
		fmt.Printf("  URL: %v\n", feed.Url)
		fmt.Print("  username: ")
		if feed.Username.Valid {
			fmt.Printf("%v", feed.Username.String)
		} else {
			fmt.Print("None")
		}
		fmt.Println()
		fmt.Println("}")
	}

	return nil
}
