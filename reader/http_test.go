package reader

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
)

var settings *viper.Viper

var _ = Describe("HTTP methods", func() {
	BeforeSuite(func() {
		viper.SetConfigName(".env.test")
		viper.AddConfigPath(".")
		viper.AddConfigPath("..")
		err := viper.ReadInConfig()

		Expect(err).To(BeNil())
		settings = viper.GetViper()
		Expect(len(settings.GetStringSlice("urls"))).To(Not(BeEquivalentTo(0)))
		logrus.SetLevel(logrus.FatalLevel)
	})

	It("Request func", func() {
		urls := settings.GetStringSlice("urls")
		for _, url := range urls {
			resp, err := request(url)
			Expect(err).To(BeNil())

			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).To(BeNil())

			err = resp.Body.Close()

			Expect(err).To(BeNil())
			Expect(len(body)).To(Not(BeEquivalentTo(0)))
		}
	})

	It("Request func with messed up url", func() {
		_, err := request("https://www.qweqweqwrvasdf.com/business-e?format=rss")
		Expect(err).To(Not(BeNil()))
	})
})
