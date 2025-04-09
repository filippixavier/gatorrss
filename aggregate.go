package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/filippixavier/gatorrss/internal/database"
	"github.com/google/uuid"
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
		publishedAt := sql.NullTime{}

		if t, err := time.Parse(time.RFC1123Z, feed.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		if _,
			err := s.db.CreatePost(context.Background(), database.CreatePostParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Title: feed.Title, Description: sql.NullString{String: feed.Description, Valid: true}, Url: feed.Link, PublishedAt: publishedAt, FeedID: source.ID}); err != nil {
			if strings.Contains(err.Error(), "posts_url_key") {
				continue
			}
			fmt.Println(err)
		}
	}

	return nil
}
