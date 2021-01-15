package main

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github-weekly-report/github"
)

var (
	name string
	message string
	messageTotal string
	star int
	fork int
	issue int
	forkPerc float64
	starPerc float64
	issuePerc float64

	initialRepo []*github.Repository 
	observerRepo []*github.Repository
)
func ReturnPercentage(beginning int, finishing int) float64 {
	if beginning == 0 {
		return 0
	}
	return float64((finishing - beginning) * 100 / beginning) 
}

func RunOnce(sizeOfRepos int, arrayofRepos []string) []*github.Repository{
	
	initialRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		initialRepo = append(initialRepo, github.InitialRepository(name))
	}

	return initialRepo
}
func RunPeroidically(sizeOfRepos int, arrayofRepos []string) []*github.Repository{

	observerRepo = nil
	for i :=0; i < sizeOfRepos ; i ++ {
		name := arrayofRepos[i]
		observerRepo = append(observerRepo, github.ObserverRepository(name))
	}

	return observerRepo
}

func Diff(i []*github.Repository, o []*github.Repository ){

	arg, _ := github.ParseArgs()
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
func SendPercentage(i []*github.Repository, o []*github.Repository)string{

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
	messageTotal = fmt.Sprintf("Changes has occurred on your repositories. \n %s ", messageTotal)
	
	return messageTotal
}
