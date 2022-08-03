package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"store/model"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (db *gorm.DB) {

	//Set Data source name
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	if os.Getenv("DB_DATABASE") == "" {
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?&parseTime=True&loc=Local",
			viper.GetString("db2.user"),
			viper.GetString("db2.pass"),
			viper.GetString("db2.host"),
			viper.GetString("db2.port"),
			viper.GetString("db2.database"),
		)
	}

	dial := mysql.Open(dsn)
	database, err := gorm.Open(dial, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("Failed to connect to database!")
	}
	//auto migration
	database.AutoMigrate(&model.Store{})
	return database
}

func Readline() (secret string, token string) {
	if os.Getenv("CHANNEL_SECRET") == "" {
		secret := viper.GetString("line.CHANNEL_SECRET")
		token := viper.GetString("line.CHANNEL_TOKEN")
		return secret, token
	} else {
		secret := os.Getenv("CHANNEL_SECRET")
		token := os.Getenv("CHANNEL_TOKEN")
		return secret, token
	}
}

func GetBot() (bot *linebot.Client) {
	secret, token := Readline()
	bot, err := linebot.New(secret, token)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func ReadMenu() (rid1 string, rid2 string) {
	if os.Getenv("CHANNEL_SECRET") == "" {
		rid1 := viper.GetString("richmenu.RID1")
		rid2 := viper.GetString("richmenu.RID2")
		return rid1, rid2
	} else {
		rid1 := os.Getenv("RID1")
		rid2 := os.Getenv("RID2")
		return rid1, rid2
	}
}

func initConfig() {
	//set Read form config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	//set timezone thailand
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func InitAll() {
	initTimeZone()
	initConfig()
}

var (
	ErrService = errors.New("service error")

	MenuFlex = `{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "https://www.img.in.th/images/a4609b9e99ef2a28a9c55ebd26935b34.png",
		  "size": "full",
		  "aspectRatio": "24:12",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"label": "Action",
			"uri": "https://linecorp.com/"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "horizontal",
		  "spacing": "md",
		  "contents": [
			{
			  "type": "box",
			  "layout": "vertical",
			  "flex": 1,
			  "margin": "xs",
			  "contents": [
				{
				  "type": "image",
				  "url": "https://www.img.in.th/images/316d54a84bd0be0e0447dbb7e75dade7.png",
				  "size": "sm",
				  "aspectRatio": "4:3"
				},
				{
				  "type": "image",
				  "url": "https://www.img.in.th/images/889028fcf2d1509b15db5424220bd8c8.png",
				  "margin": "md",
				  "size": "sm",
				  "aspectRatio": "4:3"
				},
				{
				  "type": "image",
				  "url": "https://www.img.in.th/images/896c92859ab109e7de86954e8fccf68c.png",
				  "margin": "md",
				  "size": "sm",
				  "aspectRatio": "4:3"
				},
				{
				  "type": "image",
				  "url": "https://www.img.in.th/images/a11597730e66ab8073de518442ee0983.png",
				  "margin": "md",
				  "size": "sm",
				  "aspectRatio": "4:3"
				},
				{
				  "type": "image",
				  "url": "https://www.img.in.th/images/73a965e125e523fae04685fc2569feb2.png",
				  "margin": "md",
				  "size": "sm",
				  "aspectRatio": "4:3"
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "flex": 2,
			  "spacing": "lg",
			  "contents": [
				{
				  "type": "text",
				  "text": "RELX",
				  "weight": "regular",
				  "size": "xl",
				  "flex": 2,
				  "align": "center",
				  "gravity": "center",
				  "action": {
					"type": "message",
					"text": "RELX"
				  },
				  "contents": []
				},
				{
				  "type": "separator"
				},
				{
				  "type": "text",
				  "text": "INFY",
				  "weight": "regular",
				  "size": "xl",
				  "flex": 2,
				  "align": "center",
				  "gravity": "center",
				  "action": {
					"type": "message",
					"text": "INFY"
				  },
				  "contents": []
				},
				{
				  "type": "separator"
				},
				{
				  "type": "text",
				  "text": "JUES",
				  "size": "xl",
				  "flex": 2,
				  "align": "center",
				  "gravity": "center",
				  "action": {
					"type": "message",
					"text": "JUES"
				  },
				  "contents": []
				},
				{
				  "type": "separator"
				},
				{
				  "type": "text",
				  "text": "7-11",
				  "size": "xl",
				  "flex": 2,
				  "align": "center",
				  "gravity": "center",
				  "action": {
					"type": "message",
					"text": "7-11"
				  },
				  "contents": []
				},
				{
				  "type": "separator"
				},
				{
				  "type": "text",
				  "text": "BOLD",
				  "weight": "regular",
				  "size": "xl",
				  "flex": 2,
				  "align": "center",
				  "gravity": "center",
				  "action": {
					"type": "message",
					"text": "BOLD"
				  },
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "See all",
				"text": "All Flavor"
			  }
			}
		  ]
		}
	  }`
)
