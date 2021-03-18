package github

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	owner      string
	repository string
	token      string
	fork       bool
	interval   string
)

// Flags() gets --interval argument and return th
func Flags() (int, time.Duration, bool) {

	interval = os.Getenv("INTERVAL")
	forkFlag := os.Getenv("FORK")
	if forkFlag == "" {
		fork = false
	} else {
		fork = true

	}

	// this stage control the interval argument by seperating
	// returns time choice(day or hour) and count of the time
	// duration is count of the time
	if interval == "" {
		return 0, time.Hour, fork
	} else {
		switch interval[1:] {
		case "d":
			duration, err := strconv.Atoi(string(interval[0]))
			if err != nil {
				fmt.Println(err)
			}
			return duration, 24 * time.Hour, fork

		case "h":
			duration, err := strconv.Atoi(string(interval[0]))
			if err != nil {
				fmt.Println(err)
			}
			return duration, time.Hour, fork

		case "m":
			duration, err := strconv.Atoi(string(interval[0]))
			if err != nil {
				fmt.Println(err)
			}
			return duration, time.Minute, fork
		default:
			return 0, time.Hour, fork

		}
	}
}

// ParseArgs() gets environment variables
func ParseArgs() map[string]string {

	user := make(map[string]string)

	token = os.Getenv("ACCESS_TOKEN")
	if token == "" {
		log.Fatal("Github access token is not specified")
	}
	user["token"] = token

	owner = os.Getenv("OWNER")
	if owner == "" {
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
	if repository == "" {
		log.Fatal("Repository is not specified")
	}
	user["repository"] = repository

	return user
}
