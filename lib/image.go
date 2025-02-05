package lib

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/nfnt/resize"
)

func ResizeImage(source []byte, width uint, height uint, quality int) []byte {
	bin, _, err := image.Decode(bytes.NewReader(source))
	if err != nil {
		return nil
	}
	binResized := resize.Thumbnail(width, height, bin, resize.Lanczos3)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, binResized, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil
	}
	return buf.Bytes()
}
