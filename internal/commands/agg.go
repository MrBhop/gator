package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/MrBhop/BlogAggregator/internal/database"
	"github.com/MrBhop/BlogAggregator/internal/rssFeed"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Usage: %v <time between reqs>\n", cmd.name)
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse 'time between reqs': %w\n", err)
	}
	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <- ticker.C {
		if err:= scrapeFeeds(s); err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.DataBase.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feed, err := fetchFeed(s, nextFeed)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("    - %v\n", item.Title)
	}

	return nil
}

func fetchFeed(s *state, nextFeed database.GetNextFeedToFetchRow) (*rssFeed.RSSFeed, error) {
	if err := s.DataBase.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		UpdatedAt: time.Now(),
	}); err != nil {
		return nil, err
	}

	return rssFeed.FetchFeed(context.Background(), nextFeed.Url)
}
