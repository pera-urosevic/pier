package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"pier/api/subtler/extract"

	"github.com/gin-gonic/gin"
)

func Extract(r *gin.Engine) {
	r.GET("/subtler/extract/*video", func(c *gin.Context) {
		video := os.Getenv("SUBTLER_ROOT")
		video = filepath.Join(video, c.Param("video"))

		log, err := extract.Extract(video)
		if err != nil {
			payload := map[string]string{"error": err.Error(), "log": log}
			c.JSON(http.StatusInternalServerError, payload)
			return
		}

		c.JSON(http.StatusOK, log)
	})
}
