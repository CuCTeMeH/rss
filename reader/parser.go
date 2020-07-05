package reader

import (
	"errors"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Parse exported method used externally to parse RSS feeds.
func Parse(urls []string) ([]RssItem, error) {
	if len(urls) == 0 {
		err := errors.New("no URLs to crawl")
		return []RssItem{}, err
	}

	//Make slice of urls unique.
	urls = uniqueNonEmptySlice(urls)

	//Using buffered channel to fetch data from the goroutines.
	c := make(chan *gofeed.Feed, len(urls))

	//Fetch the feeds concurrently.
	for _, url := range urls {
		go read(url, c)
	}

	var result []RssItem

	//Loop and set the type we need to return.
	for i := 0; i < len(urls); i++ {
		chanRes := <-c
		if chanRes != nil {
			for _, v := range chanRes.Items {
				item := RssItem{
					Source:      chanRes.Title,
					SourceURL:   chanRes.Link,
					Title:       v.Title,
					Description: v.Description,
					Link:        v.Link,
					PublishDate: *v.PublishedParsed,
				}
				result = append(result, item)
			}
		}
	}

	return result, nil
}

//Parse the request from the url with gofeed library. Will parse Atom and RSS feeds.
func parseRequest(resp *http.Response) *gofeed.Feed {
	defer resp.Body.Close()
	parser := gofeed.NewParser()
	feed, err := parser.Parse(resp.Body)
	if err != nil {
		logrus.WithError(err)
		return nil
	}
	return feed
}

//Make a slice unique, using this to make sure that we crawl only once for a given url.
func uniqueNonEmptySlice(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us
}
