package main

import (
	"time"

	"github.com/Prakhar256/RSS_aggregator/internal/database"
	"github.com/google/uuid"
)
type User struct{
	ID uuid.UUID   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	APIKey string `json:"api_key"`
}
func databaseUserToUser(dbUser database.User) User{
	return User{
		ID: dbUser.ID,
		CreatedAt:dbUser.CreatedAt,
		UpdatedAt:dbUser.UpdatedAt,
		Name:dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}  
type Feed struct{
	ID uuid.UUID   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name string `json:"name"`
	Url string	`json:"url"` 
	UserId uuid.UUID  `json:"user_id"`
}
func databaseFeedToFeed(dbFeed database.Feed) Feed{
	return Feed{
		ID: dbFeed.ID,
		CreatedAt:dbFeed.CreatedAt,
		UpdatedAt:dbFeed.UpdatedAt,
		Name:dbFeed.Name,
		Url:dbFeed.Url,
		UserId: dbFeed.ID,
	}
}  
func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed{
	feeds:=[]Feed{}
	for _, value:=range dbFeed{
		feeds=append(feeds, databaseFeedToFeed(value))
	}
	return feeds
}  