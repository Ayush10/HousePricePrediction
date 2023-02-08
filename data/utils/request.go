package utils

import (
	"io/ioutil"
	"net/http"
)

// GetPage retrieves the HTML content of a given URL
func GetPage(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return html, nil
}
