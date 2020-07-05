package reader

import "time"

type RssItem struct {
	Source      string
	SourceURL   string
	Title       string
	Description string
	Link        string
	PublishDate time.Time
}
