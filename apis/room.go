package apis

import (
	"admin-go/template"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{} // use default options

func Server(c *gin.Context) {
	server, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("upgrade errors, errors msg are: %v", err.Error())
		return
	}
	defer server.Close()
	for {
		messageType, message, err := server.ReadMessage()
		if err != nil {
			log.Printf("read errors, errors msg are: %v", err.Error())
			break
		}
		log.Printf("recv msg: %v", string(message))
		err = server.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("write errors, errors msg are: %v", err.Error())
			break
		}
	}
}

func Room(c *gin.Context)  {
	template.Home.Execute(c.Writer,"ws://"+c.Request.Host+"/socket/server")
}
