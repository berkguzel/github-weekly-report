package main

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/berkguzel/github-weekly-report/github"
)

var (
	name string
	star int
	fork int
	issue int
	forkPerc float64
	starPerc float64
	issuePerc float64

)

// ReturnPercentage() calculate the differences
func ReturnPercentage(beginning int, finishing int) float64 {

	// To avoid "zero value" error
	if beginning == 0 {
		return float64((finishing - beginning) * 100 / 1)
	}
	return float64((finishing - beginning) * 100 / beginning) 
}

// Diff() make comparison between InitialRepoistory and ObserverRepository and sends message
func Diff(i []github.Repository, o []github.Repository ){
	
	var message string
	arg := github.ParseArgs()
	bot, err := tgbotapi.NewBotAPI(arg["tgBotToken"])
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	chatId, _ := strconv.ParseInt(arg["chatId"], 10, 64)

	for r , v := range o {
		name = v.Name
		star = v.StargazersCount - i[r].StargazersCount
		fork = v.ForksCount - i[r].ForksCount
		issue = v.OpenIssuesCount - i[r].OpenIssuesCount

		message = message + "\n" + fmt.Sprintf("Repository: %s \n New Stars: %d \n New Forks: %d \n New Issues: %d \n",
		name, star, fork, issue )
	}
	message = fmt.Sprintf("Hi, %s! Here is what occurred on your repositories. \n %s ", arg["owner"], message)
	msg := tgbotapi.NewMessage(chatId, message)
	bot.Send(msg)

	message = SendPercentage(i, o)
	msg = tgbotapi.NewMessage(chatId, message)
	bot.Send(msg)	
}

// SendPercentage() sends a new message about change rate 
func SendPercentage(i []github.Repository, o []github.Repository)string{

	var messageTotal string
	for r , v := range o {
		name = v.Name
		fork = v.ForksCount
		forkPerc = ReturnPercentage(i[r].ForksCount, v.ForksCount)
		star = v.StargazersCount
		starPerc = ReturnPercentage(i[r].StargazersCount, v.StargazersCount)
		issue = v.OpenIssuesCount
		issuePerc = ReturnPercentage(i[r].OpenIssuesCount, v.OpenIssuesCount )
		messageTotal = messageTotal + "\n" + fmt.Sprintf("Repository: %s \n Star Count: %d  Change: %s %.2f \n Fork Count: %d Change: %s %.2f \n Issue Count: %d Change: %s %.2f \n", name, star, "%", starPerc, fork, "%", forkPerc, issue,"%", issuePerc )
	}
	messageTotal = fmt.Sprintf("Changes have occurred on your repositories. \n %s ", messageTotal)
	
	return messageTotal
}
