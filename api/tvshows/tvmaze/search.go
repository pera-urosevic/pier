package tvmaze

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Search(title string) (string, error) {
	var results string

	url := fmt.Sprintf("%s/search/shows?q=%s", site, url.QueryEscape(title))
	res, err := http.Get(url)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return results, err
	}

	results = string(bytes)
	return results, nil
}
