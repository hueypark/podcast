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
	Items       []Item
}

func MakeFeed(feed *gofeed.Feed) Feed {
	items := make([]Item, len(feed.Items))
	for i, item := range feed.Items {
		items[i] = MakeItem(item)
	}

	return Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		FeedLink:    feed.FeedLink,
		Updated:     *feed.UpdatedParsed,
		Items:       items,
	}
}
