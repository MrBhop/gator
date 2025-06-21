package rssFeed

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}


func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Printf("fetching from %v ...\n", feedUrl)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var output RSSFeed
	if err := xml.Unmarshal(content, &output); err != nil {
		return nil, err
	}

	output = unescapeHtml(output)

	return &output, nil
}

func unescapeHtml(intput RSSFeed) RSSFeed {
	output := RSSFeed{}

	output.Channel.Title = html.UnescapeString(intput.Channel.Title)
	output.Channel.Description = html.UnescapeString(intput.Channel.Description)
	
	for _, item := range intput.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)

		output.Channel.Item = append(output.Channel.Item, item)
	}

	return output
}
