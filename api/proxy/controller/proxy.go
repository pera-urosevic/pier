package controller

import (
	"bytes"
	"io"
	"net/http"
	"pier/api/proxy/types"
	"pier/lib"

	"github.com/gin-gonic/gin"
)

func Proxy(r *gin.Engine) {
	r.POST("/proxy", func(c *gin.Context) {
		var params types.ProxyParams

		err := c.BindJSON(&params)
		if err != nil {
			lib.GinError(c, err, false)
		}

		reqBody := bytes.NewBuffer([]byte(params.Body))
		req, err := http.NewRequest(params.Method, params.URL, reqBody)
		if err != nil {
			lib.GinError(c, err, true)
		}

		for key, value := range params.Headers {
			req.Header.Add(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			lib.GinError(c, err, true)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			lib.GinError(c, err, true)
		}
		result := string(body)

		c.JSON(res.StatusCode, result)
	})
}
