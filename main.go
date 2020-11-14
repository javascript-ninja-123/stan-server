package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	listener "github.com/javascript-ninja-123/stan-server/listener"
	subject "github.com/javascript-ninja-123/stan-server/subject"
	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("production", "randomId")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	onMessage := func(msg *stan.Msg) {
		fmt.Println(string(msg.Data))
		msg.Ack()
	}
	entityUploadListener := listener.NewBaseListener(subject.EntityUploaded, "coolqueueName", time.Second*5, onMessage, sc)
	entityUploadListener.Listen()

	r.Run()
}
