package apis

import (
	"admin-go/defs"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
)

func VideoApi(c *gin.Context) {
	video, err := ioutil.ReadFile(defs.BASE_VIDEO_PATH + c.Query("video-id"))

	if err != nil {
		log.Printf("sorry open video error %v", err.Error())
		c.Abort()
		return
	}
	//c.SSEvent("video file", video)
	c.Stream(func(w io.Writer) bool {
		w.Write(video)
		return true
	})
	//c.Data(http.StatusOK, "video/mp4", video)
}
