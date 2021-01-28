package github

import (
	"flag"
	"os"
	"log"
	"time"
	"strconv"
	"fmt"
)
var (
	owner string
	repository string
	token string
	duration int
	interval *string
	fork *bool
	intervalSlice string
	forkFlag bool
)

// Flags() gets --interval argument and return th 
func Flags()(int, time.Duration, bool){
	
	interval = flag.String("interval", "", "interval to notify")
	fork = flag.Bool("fork", false, "true returns the forked repositories")
	flag.Parse()
	intervalSlice = *interval
	forkFlag = *fork

	// this stage control the interval argument by seperating
	// returns time choice(day or hour) and count of the time
	// duration is count of the time
	if intervalSlice == "" {
		return 0, time.Hour, forkFlag
	} else {
		switch intervalSlice[1:] {
		case "d":
			duration, err := strconv.Atoi(string(intervalSlice[0]))
			if err != nil {
				fmt.Println(err)
			}
			return duration, 24* time.Hour, forkFlag

		case "h":
			duration, err := strconv.Atoi(string(intervalSlice[0]))
			if err != nil {
				fmt.Println(err)
			}
			return duration, time.Hour, forkFlag
		
		case "m":
			duration, err := strconv.Atoi(string(intervalSlice[0]))
			if err != nil{
				fmt.Println(err)
			}
			return duration, time.Minute, forkFlag

		default :
			return 0, time.Hour, forkFlag
					
		}
	}
}

// ParseArgs() gets environment variables
func ParseArgs() map[string]string {	
	
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

	return user
}