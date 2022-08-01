package handler

import (
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
				if message.Text == "ยกเลิกคิว" {
					// good, err := h.qService.DeleteGoodbyUID(event.Source.UserID)
					if err != nil {
						if err.Error() == "user Code not found" {
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("ท่านยังไม่ได้จองคิวไม่สามารถยกเลิกได้")).Do(); err != nil {
								log.Print(err)
							}
							return
						} else {
							if _, err := bot.PushMessage(event.Source.UserID, linebot.NewTextMessage("เกิดข้อผิดพลาดไม่สามารถยกเลิกคิวได้")).Do(); err != nil {
								log.Print(err)
							}
							return
						}
					}
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("ท่านยกเลิกคิว %v เรียบร้อยแล้ว", good.Code))).Do(); err != nil {
					// log.Print(err)
					// }
					return
				}

				// split := strings.Split(message.Text, " ")
				// if split[0] == "ดู" || split[0] == "ตรวจสอบ" || split[0] == "ค้นหา" {
				// 	fmt.Println(split[1])
				// 	flex, err := h.qService.FlexGood(split[1])
				// 	fmt.Println(err)
				// 	if err != nil {
				// 		if err.Error() == "repository error" {
				// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ไม่พบเลขคิวที่คุณค้นหาหรืออาจเลยคิวของคุณมาแล้ว")).Do(); err != nil {
				// 				log.Print(err)
				// 				return
				// 			}
				// 			return
				// 		} else {
				// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ระบบผิดพลาด")).Do(); err != nil {
				// 				log.Print(err)
				// 				return
				// 			}
				// 			return
				// 		}
				// 	}
				// 	// Unmarshal JSON
				// 	flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(flex))
				// 	if err != nil {
				// 		log.Println(err)
				// 	}
				// 	// New Flex Message
				// 	flexMessage := linebot.NewFlexMessage(split[1], flexContainer)
				// 	// Reply Message
				// 	_, err = bot.ReplyMessage(event.ReplyToken, flexMessage).Do()
				// 	if err != nil {
				// 		log.Print(err)
				// 	}
				// } else {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ขออภัยครับ แต่เรายังไม่เข้าใจ ท่านอยากจะทวนอีกรอบหรือส่งต่อให้เจ้าหน้าที่ตอบคำถามดีครับ")).Do(); err != nil {
					log.Print(err)
				}
				// }
			}
		}
	}
}
