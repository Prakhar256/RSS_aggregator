package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/Prakhar256/RSS_aggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", timeBetweenRequest, concurrency)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Couldn't get next feeds to fetch", err)
			continue
		}
		log.Printf("Found %v feeds to fetch!", len(feeds))

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
    // Parse the PubDate string into a time.Time value
    pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
    if err != nil {
        log.Printf("Couldn't parse publication date for post %s: %v", item.PubDate, err)
        continue // Skip this item if the date can't be parsed
    }

    // Insert the post into the database
    _,err = db.CreatePost(context.Background(), database.CreatePostParams{
        ID:          uuid.New(),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        Title:       item.Title,
        Description: sql.NullString{
            String: item.Description,
            Valid:  item.Description != "",
        },
        PublishedAt: pubDate,  // Use the parsed pubDate here
        Url:         item.Link,
        FeedID:      feed.ID,
    })

    if err != nil {
        log.Printf("Couldn't create post for feed %s: %v", feed.Name, err)
    }
}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
