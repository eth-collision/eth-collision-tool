package tool

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var config Config
var bot *tgbotapi.BotAPI

func init() {
	config.Init()
}

type Config struct {
	Token  string `yaml:"token"`
	ChatId int64  `yaml:"chat-id"`
}

func (c *Config) Init() {
	yamlFile, err := ioutil.ReadFile("telegram-config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
		return
	}
	b, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Println(err)
		return
	}
	bot = b
}

func SendMsgText(text string) {
	if bot == nil {
		log.Println("bot is nil")
		return
	}
	message := tgbotapi.NewMessage(config.ChatId, text)
	_, err := bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}
