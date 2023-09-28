package feed

import "github.com/mmcdole/gofeed"

type Item struct {
	GUID        string
	Title       string
	Description string
	Updated     string
	Enclosures  []Enclosure
}

func MakeItem(item *gofeed.Item) Item {
	enclosures := make([]Enclosure, len(item.Enclosures))
	for i, enclosure := range item.Enclosures {
		enclosures[i] = MakeEnclosure(enclosure)
	}

	return Item{
		GUID:        item.GUID,
		Title:       item.Title,
		Description: item.Description,
		Updated:     item.Updated,
		Enclosures:  enclosures,
	}
}
