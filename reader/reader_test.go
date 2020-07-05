package reader

import (
	"github.com/mmcdole/gofeed"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reader methods", func() {

	It("Read Feed via HTTP", func() {
		urls := settings.GetStringSlice("urls")
		c := make(chan *gofeed.Feed, len(urls))

		for _, url := range urls {
			go read(url, c)
		}

		var result []RssItem

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

		Expect(result).To(Not(BeNil()))
	})

	It("Not correct url for request", func() {
		urls := settings.GetStringSlice("urls")
		for key, url := range urls {
			urls[key] = url + "123"
		}

		c := make(chan *gofeed.Feed, len(urls))

		for _, url := range urls {
			go read(url, c)
		}

		var result []RssItem

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

		Expect(result).To(BeNil())
	})

	It("Empty url for request", func() {
		urls := []string{}
		c := make(chan *gofeed.Feed, len(urls))

		for _, url := range urls {
			go read(url, c)
		}

		var result []RssItem

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

		Expect(result).To(BeNil())
	})

	It("404 url for request", func() {
		urls := []string{
			"httpws://en.wikipedia.org/page",
		}
		c := make(chan *gofeed.Feed, len(urls))

		for _, url := range urls {
			go read(url, c)
		}

		var result []RssItem

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

		Expect(result).To(BeNil())
	})
})
