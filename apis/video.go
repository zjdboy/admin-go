package apis

import (
	"admin-go/defs"
	"admin-go/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func VideoApi(c *gin.Context) {
	video, err := ioutil.ReadFile(defs.BASE_VIDEO_PATH + c.Query("video-id"))

	if err != nil {
		log.Printf("sorry open video error %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Reached the limit connections",
		})
		return
	}
	//c.SSEvent("video file", video)
	c.Stream(func(w io.Writer) bool {
		limiter := middleware.NewLimit(3)
		_, err := w.Write(video)

		if err != nil {
			defer limiter.Release()
		}
		return true
	})
	//c.Data(http.StatusOK, "video/mp4", video)
}
