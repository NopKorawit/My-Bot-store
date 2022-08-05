package handler

import (
	"Product/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func (h ProductHandler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func (h ProductHandler) Callback(c *gin.Context) {
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
				if message.Text == "RELX" || message.Text == "INFY" || message.Text == "JUES" || message.Text == "7-11" || message.Text == "BOLD" || message.Text == "See all" {
					var types string
					switch message.Text {
					case "See all":
						// types = "All"
					case "RELX":
						types = "A"
					case "INFY":
						types = "B"
					case "JUES":
						types = "C"
					case "7-11":
						types = "D"
					case "BOLD":
						types = "D"
					default:
						log.Println("This Type not in Conditions")
					}

					Products, err := h.qService.GetProductsType(types)
					if err != nil {
						if err.Error() == "queue already exists" {
							if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ท่านจองคิวไปแล้วกรุณายกเลิกคิวก่อนหน้า")).Do(); err != nil {
								log.Print(err)
							}
							return
						} else {
							if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาดไม่สามารถบันทึกคิวได้")).Do(); err != nil {
								log.Print(err)
							}
							return
						}
					}

					head := fmt.Sprintf("รายการ %v ตามนี้ค้าบ\n", message.Text)
					var quantity string
					for _, Product := range Products {

						if Product.Quantity == 0 {
							quantity = "❌"
						} else if Product.Quantity < 3 {
							quantity = "⚠️"
						} else {
							quantity = "✅"
						}
						text := fmt.Sprintf("%v | %v | %v\n", quantity, Product.Code, Product.Name)
						head = head + text
					}
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(head)).Do(); err != nil {
						log.Print(err)
					}
					return
				}
				rows := strings.Split(message.Text, "\n")
				if rows[0] == "ซื้อ" || rows[0] == "เอา" || rows[0] == "buy" || rows[0] == "order" {
					rows := rows[1:]
					text := "ซื้อ\n"
					for _, row := range rows {
						split := strings.Split(row, " ")
						fmt.Println(split[1])
						amount, _ := strconv.Atoi(split[1])
						Product, err := h.qService.SellProduct(split[0], amount)
						fmt.Println(err)
						if err != nil {
							if err == model.ErrProductNotEnough {
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("สินค้ามีจำนวนไม่เพียงพอ")).Do(); err != nil {
									log.Print(err)
									return
								}
								return
							} else if err == model.ErrCodenotFound {
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ค้นหาสินค้าไม่เจอ")).Do(); err != nil {
									log.Print(err)
									return
								}
								return
							} else {
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ระบบผิดพลาด")).Do(); err != nil {
									log.Print(err)
									return
								}
								return
							}
						}
						list := fmt.Sprintf("%v หัว| %v %v\n", Product.Quantity, Product.Type, Product.Name)
						text = text + list
						fmt.Println(Product)
					}
					text = text + "เรียบร้อยแล้ว"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
						log.Print(err)
						return
					}
				} else {
					// Emoji
					sorry := linebot.NewEmoji(0, "5ac1bfd5040ab15980c9b435", "024")
					// have := linebot.NewEmoji(0, "5ac21a18040ab15980c9b43e", "007")
					// out := linebot.NewEmoji(0, "5ac21a18040ab15980c9b43e", "068")
					// few := linebot.NewEmoji(0, "5ac21a18040ab15980c9b43e", "025")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("$ ขออภัยครับ แต่ผมยังไม่เข้าใจ ท่านอยากจะทวนอีกรอบหรือรอให้นพมาตอบคำถามดีครับ").AddEmoji(sorry)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
