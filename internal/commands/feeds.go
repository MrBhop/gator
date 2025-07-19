package commands

import (
	"context"
	"fmt"

	"github.com/MrBhop/gator/internal/database"
)

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.DataBase.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(GetFeedsRowToString(feed))
	}

	return nil
}

func GetFeedsRowToString(item database.GetFeedsRow) string {
	output := "GetFeedsRow{"
	output += fmt.Sprintf("Name: %v\n", item.Name)
	output += fmt.Sprintf("URL: %v\n", item.Url)
	output += "Username: "
	if item.Username.Valid {
		output += fmt.Sprint(item.Username)
	} else {
		output += "None"
	}
	output += fmt.Sprintln()
	output += fmt.Sprintln("}")
	return output
}
