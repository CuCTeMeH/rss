package reader

import (
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

//HTTP request for the url and return the feed via channel.
func read(url string, c chan *gofeed.Feed) {
	resp, err := request(url)

	if err != nil {
		c <- nil
		return
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			c <- nil
			return
		}

		bodyString := string(bodyBytes)
		logrus.WithField("body", bodyString).Error(resp.Status)
		c <- nil
		return
	}

	c <- parseRequest(resp)
}
