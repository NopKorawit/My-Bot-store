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
		  "url": "https://www.i-pic.info/i/KMdp196143.png",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover"
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "Booking",
			  "weight": "bold",
			  "size": "xl",
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "baseline",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "text",
				  "text": "How many people are you?",
				  "size": "lg",
				  "color": "#AAAAAA",
				  "flex": 1,
				  "contents": []
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "flex": 0,
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "Alone",
				"text": "Alone"
			  },
			  "height": "40px",
			  "backgroundColor": "#FBF1C2FF",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "Alone",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "Couple",
				"text": "Couple"
			  },
			  "height": "40px",
			  "backgroundColor": "#DFE9F5",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "Couple",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "Small Group (3-4)",
				"text": "Small Group"
			  },
			  "height": "40px",
			  "backgroundColor": "#F6E6DE",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "Small Group (3-4)",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "action": {
				"type": "message",
				"label": "The Gang (5-6)",
				"text": "The Gang"
			  },
			  "height": "40px",
			  "backgroundColor": "#D9F0E7",
			  "cornerRadius": "8px",
			  "contents": [
				{
				  "type": "text",
				  "text": "The Gang (5-6)",
				  "color": "#000000FF",
				  "align": "center",
				  "margin": "md",
				  "contents": []
				}
			  ]
			},
			{
			  "type": "spacer"
			}
		  ]
		}
	  }`
)
