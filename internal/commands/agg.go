package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MrBhop/BlogAggregator/internal/database"
	"github.com/MrBhop/BlogAggregator/internal/rssFeed"
	"github.com/google/uuid"
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
		pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			return fmt.Errorf("couldn't parse pubDate: %w\n", err)
		}

		newPost, err := s.DataBase.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID: nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			return fmt.Errorf("couldn't create post: %w\n", err)
		}

		fmt.Printf("    created post: %+v\n", newPost)
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
