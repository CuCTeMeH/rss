package reader

import (
	"net/http"
)

//Make a request for given url.
func request(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, err
}
