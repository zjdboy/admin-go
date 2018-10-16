package handlers

import (
	"admin-go/defs"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func VideoHandler(c *gin.Context) {
	video, err := ioutil.ReadFile(defs.BASE_VIDEO_PATH + c.Query("video-id"))

	if err != nil {
		log.Fatalf("sorry open video error %v", err.Error())
	}

	c.Data(http.StatusOK, "video/mp4", video)
}
