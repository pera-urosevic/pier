package tvmaze

import (
	"fmt"
	"io"
	"net/http"
)

func Get(id int64) (string, error) {
	var result string

	url := fmt.Sprintf("%s/shows/%d?embed=episodes", site, id)
	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	result = string(bytes)
	return result, nil
}
