package feed

import (
	"time"

	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Title       string
	Description string
	Link        string
	FeedLink    string
	Updated     time.Time
	Episodes    []Episode
}

func MakeFeed(feed *gofeed.Feed) Feed {
	episodes := make([]Episode, len(feed.Items))
	for i, item := range feed.Items {
		episodes[i] = MakeEpisode(item)
	}

	return Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		FeedLink:    feed.FeedLink,
		Updated:     *feed.UpdatedParsed,
		Episodes:    episodes,
	}
}
