package github

import (
	"os"
	"log"
)
var (
	owner string
	repository string
	token string
	fork bool
	notify string
	difference bool
	inerval string
)

func ParseArgs() (map[string]string, bool) {	
	
	user := make(map[string]string)

	token = os.Getenv("ACCESS_TOKEN")
	if token == ""{
		log.Fatal("Github access token is not specified")
	}
	user["token"] = token
	
	owner = os.Getenv("OWNER")
	if owner == ""{
		log.Fatal("Owner of the repositories is not specified")
	}
	user["owner"] = owner
	
	chatId := os.Getenv("CHATID")
	if chatId == "" {
		log.Fatal("Chat Id is not specified")
	}
	user["chatId"] = chatId

	tgBotToken := os.Getenv("TOKEN")
	if tgBotToken == "" {
		log.Fatal("Telegram Bot token is not specified")
	}
	user["tgBotToken"] = tgBotToken

	repository = os.Getenv("REPOSITORY")
	if repository == ""{
		log.Fatal("Repository is not specified")
	}
	user["repository"] = repository
	// TODO be sure user seperated repositories with comma ,

	return user, false
}