package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/filippixavier/gatorrss/internal/database"
)

func scrapeFeeds(s *state) error {
	source, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{ID: source.ID, LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}})

	feeds, err := fetchFeed(context.Background(), source.Url)

	if err != nil {
		return err
	}

	for _, feed := range feeds.Channel.Item {
		fmt.Println(feed.Title)
	}

	return nil
}
