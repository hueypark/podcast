package feed

import "github.com/mmcdole/gofeed"

type Enclosure struct {
	URL    string
	Length string
	Type   string
}

func MakeEnclosure(enclosure *gofeed.Enclosure) Enclosure {
	return Enclosure{
		URL:    enclosure.URL,
		Length: enclosure.Length,
		Type:   enclosure.Type,
	}
}
