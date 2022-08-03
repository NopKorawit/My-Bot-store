package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func (h goodHandler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func (h goodHandler) Callback(c *gin.Context) {
	bot := GetBot()
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(400)
		} else {
			c.Writer.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "test" {

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("test")).Do(); err != nil {
						log.Print(err)
					}
				}
				if message.Text == "Flavor" {
					// Unmarshal JSON
					flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(MenuFlex))
					if err != nil {
						log.Println(err)
					}
					fmt.Println(flexContainer)
					// New Flex Message
					flexMessage := linebot.NewFlexMessage(message.Text, flexContainer)
					// Reply Message
					_, err = bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
					if err != nil {
						log.Print(err)
					}
					return
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ขออภัยครับ แต่เรายังไม่เข้าใจ ท่านอยากจะทวนอีกรอบหรือส่งต่อให้เจ้าหน้าที่ตอบคำถามดีครับ")).Do(); err != nil {
					log.Print(err)
				}
				// }
			}
		}
	}
}
