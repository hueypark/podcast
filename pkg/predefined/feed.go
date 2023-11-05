package predefined

// Feeds is a list of predefined feeds.
var Feeds = []Feed{
	{
		Title: "Cup o' Go",
		Link:  "https://feeds.transistor.fm/cup-o-go",
	},
	{
		Title: "Go Time: Golang, Software Engineering",
		Link:  "https://changelog.com/gotime/feed",
	},
	{
		Title: "Lex Fridman Podcast",
		Link:  "https://lexfridman.com/feed/podcast/",
	},
	{
		Title: "All-In with Chamath, Jason, Sacks &amp; Friedberg",
		Link:  "https://feeds.libsyn.com/254861/rss",
	},
	{
		Title: "This Week in Startups",
		Link:  "https://anchor.fm/s/7c624c84/podcast/rss",
	},
	{
		Title: "I Will Teach You To Be Rich",
		Link:  "https://feeds.megaphone.fm/IWTYB6912370287/",
	},
}

type Feed struct {
	Title string
	Link  string
}
