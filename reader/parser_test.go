package reader

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser methods", func() {
	It("Parse URLs to struct", func() {
		urls := settings.GetStringSlice("urls")

		rssItems, err := Parse(urls)
		Expect(err).To(BeNil())
		Expect(len(rssItems)).To(Not(BeEquivalentTo(0)))
	})

	It("Parse empty slice of URLs to struct", func() {
		urls := []string{}

		rssItems, err := Parse(urls)
		Expect(err).To(Not(BeNil()))
		Expect(len(rssItems)).To(BeEquivalentTo(0))
	})

	It("Check Unique slice func", func() {
		urls := settings.GetStringSlice("urls")
		urlsDuplicated := append(urls, urls...)

		slice := uniqueNonEmptySlice(urlsDuplicated)

		Expect(len(slice)).To(BeEquivalentTo(len(urls)))
	})
})
