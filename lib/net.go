package lib

import (
	"encoding/base64"
	"io"
	"net/http"
)

func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Base64Encode(bytes []byte) string {
	mimeType := http.DetectContentType(bytes)
	base64Encoding := "data:" + mimeType + ";base64,"
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding
}
