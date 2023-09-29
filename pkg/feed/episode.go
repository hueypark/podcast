package feed

import "github.com/mmcdole/gofeed"

type Episode struct {
	GUID        string
	Title       string
	Description string
	Updated     string
	Enclosures  []Enclosure
}

func MakeEpisode(item *gofeed.Item) Episode {
	enclosures := make([]Enclosure, len(item.Enclosures))
	for i, enclosure := range item.Enclosures {
		enclosures[i] = MakeEnclosure(enclosure)
	}

	return Episode{
		GUID:        item.GUID,
		Title:       item.Title,
		Description: item.Description,
		Updated:     item.Updated,
		Enclosures:  enclosures,
	}
}
