package commands

import (
	"context"
	"fmt"

	"github.com/MrBhop/BlogAggregator/internal/rssFeed"
)


func handlerAgg(s *state, _ command) error {
	const url string = "https://www.wagslane.dev/index.xml"

	feed, err := rssFeed.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("Feed fetched:\n%+v\n", feed)
	return nil
}
